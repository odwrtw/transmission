package transmission

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
)

var (
	defaultAddress = "http://localhost"
	defaultPort    = "9091"
	defaultPath    = "/transmission/rpc"
	defaultScheme  = "http"
)

type Config struct {
	// Address support format
	// http://localhost:9091/transmission/rcp
	// localhost
	// localhost/transmission/rcp
	// localhost:9091
	Address      string
	Port         string
	User         string
	Password     string
	skipCheckSSL bool
}

type Client struct {
	httpClient *http.Client
	conf       *Config
	sessionID  string
	endpoint   string
}

type GetTorrentArg struct {
	Fields []string `json:"fields,omitempty"`
	Ids    []int    `json:"ids,omitempty"`
}

type AddTorrentArg struct {
	// Cookies string
	// download-dir string
	// Filename filename or URL of the .torrent file
	Filename string `json:"filename,omitempty"`
	// Metainfo base64-encoded .torrent content
	Metainfo string `json:"metainfo,omitempty"`
	// Paused   bool
	// peer-limit int
	// BandwidthPriority int
	// files-wanted
	// files-unwanted
	// priority-high
	// priority-low
	// priority-normal

}

type RemoveTorrentArg struct {
	Ids             []int `json:"ids,string"`
	DeleteLocalData bool  `json:"delete-local-data,omitempty"`
}

type Request struct {
	Method    string      `json:"method"`
	Arguments interface{} `json:"arguments"`
}

type Response struct {
	Arguments interface{} `json:"arguments"`
	Result    string      `json:"result"`
}

func (c *Client) Do(req *http.Request, retry bool) (*http.Response, error) {
	if c.conf.User != "" && c.conf.Password != "" {
		req.SetBasicAuth(c.conf.User, c.conf.Password)
	}
	if c.sessionID != "" {
		req.Header.Add("X-Transmission-Session-Id", c.sessionID)
	}

	//Body copy for replay it if needed
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	req.Body = ioutil.NopCloser(bytes.NewBuffer(b))

	//Log request for debug
	log.Print(bytes.NewBuffer(b).String())

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	// error 409
	if resp.StatusCode == http.StatusConflict && retry {
		c.sessionID = resp.Header.Get("X-Transmission-Session-Id")
		req.Body = ioutil.NopCloser(bytes.NewBuffer(b))
		return c.Do(req, false)
	}
	return resp, nil
}

func (c *Client) post(tReq *Request) (*http.Response, error) {
	data, err := json.Marshal(tReq)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", c.endpoint, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	return c.Do(req, true)
}

func (c *Client) request(tReq *Request, tResp *Response) error {
	resp, err := c.post(tReq)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, tResp)
	if err != nil {
		return err
	}
	if tResp.Result != "success" {
		return fmt.Errorf("transmission: request response %q", tResp.Result)
	}
	return nil
}

func (c *Client) GetTorrents() (*[]Torrent, error) {
	tReq := &Request{
		Arguments: GetTorrentArg{
			Fields: torrentGetFields,
		},
		Method: "torrent-get",
	}

	r := &Response{Arguments: &Torrents{}}

	err := c.request(tReq, r)
	if err != nil {
		return nil, err
	}

	t := *r.Arguments.(*Torrents).Torrents
	for i := 0; i < len(t); i++ {
		t[i].Client = c
	}
	return &t, nil
}

func (c *Client) AddTorrent(filename, metadata string) (*Torrent, error) {
	tReq := &Request{
		Arguments: AddTorrentArg{
			Filename: filename,
			Metainfo: metadata,
		},
		Method: "torrent-add",
	}
	type added struct {
		Torrent *Torrent `json:"torrent-added"`
	}
	r := &Response{Arguments: &added{}}
	err := c.request(tReq, r)
	if err != nil {
		return nil, err
	}
	t := r.Arguments.(*added)
	t.Torrent.Client = c
	return t.Torrent, nil
}

func (c *Client) RemoveTorrents(torrents []*Torrent, removeData bool) error {
	ids := make([]int, len(torrents))
	for i := range torrents {
		ids[i] = torrents[i].Id
	}
	tReq := &Request{
		Arguments: RemoveTorrentArg{
			Ids:             ids,
			DeleteLocalData: removeData,
		},
		Method: "torrent-remove",
	}
	r := &Response{}
	err := c.request(tReq, r)
	if err != nil {
		return err
	}
	return nil
}

func New(conf Config) (*Client, error) {
	httpClient := &http.Client{}
	if conf.skipCheckSSL {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		httpClient = &http.Client{Transport: tr}
	}
	if conf.Address == "" {
		conf.Address = defaultAddress
	}
	if conf.Port == "" {
		conf.Port = defaultPort
	}
	u, err := url.Parse(conf.Address)
	if err != nil {
		return nil, err
	}

	// Support for address type "localhost"
	if u.Host == "" && u.Path != "" {
		u.Host = u.Path
		u.Path = ""
	}

	// Support for addres type "locahost:9091"
	if u.Scheme != "" && u.Opaque != "" && u.Host == "" {
		u.Host = u.Scheme + ":" + u.Opaque
		u.Scheme = ""
		u.Opaque = ""
	}

	if u.Scheme == "" {
		u.Scheme = defaultScheme
	}
	_, _, err = net.SplitHostPort(u.Host)
	if err != nil {
		//error if :port not set
		u.Host = u.Host + ":" + conf.Port
	}
	if u.Path == "" {
		u.Path = defaultPath
	}
	return &Client{conf: &conf, httpClient: httpClient, endpoint: u.String()}, nil
}

package transmission

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

var (
	defaultAddress = "http://localhost"
	defaultPort    = "9091"
)

type Config struct {
	address      string
	port         string
	user         string
	password     string
	skipCheckSSL bool
}

type Client struct {
	httpClient *http.Client
	conf       *Config
	sessionID  string
	endpoint   string
}

type ReqArguments struct {
	Fields   []string `json:"fields,omitempty"`
	Filename string   `json:"filename,omitempty"`
	Ids      []int    `json:"ids,omitempty"`
}

type ReqArgs struct {
	Method    string       `json:"method"`
	Arguments ReqArguments `json:"arguments"`
}

type Response struct {
	Arguments interface{} `json:"arguments"`
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	if c.conf.user != "" && c.conf.password != "" {
		req.SetBasicAuth(c.conf.user, c.conf.password)
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

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 409 {
		c.sessionID = resp.Header.Get("X-Transmission-Session-Id")
		req.Body = ioutil.NopCloser(bytes.NewBuffer(b))
		return c.Do(req)
	}
	return resp, nil
}

func (c *Client) Post(method string) (*http.Response, error) {
	postData := &ReqArgs{
		Arguments: ReqArguments{
			Fields: torrentGetFields,
		},
		Method: method}
	data, err := json.Marshal(postData)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", c.endpoint, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}

func (c *Client) GetTorrents() (*[]Torrent, error) {
	resp, err := c.Post("torrent-get")
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	r := Response{Arguments: &Torrents{}}
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}
	t := r.Arguments.(*Torrents).Torrents
	return t, nil
}

func New(conf Config) (*Client, error) {
	httpClient := &http.Client{}
	if conf.skipCheckSSL {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		httpClient = &http.Client{Transport: tr}
	}
	if conf.address == "" {
		conf.address = defaultAddress
	}
	if conf.port == "" {
		conf.port = defaultPort
	}
	u, err := url.Parse(conf.address)
	if err != nil {
		return nil, err
	}
	u.Host = u.Host + ":" + conf.port
	u.Path = "/transmission/rpc"
	return &Client{conf: &conf, httpClient: httpClient, endpoint: u.String()}, nil
}

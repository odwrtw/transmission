package transmission

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	// DefaultAddress default transmission address
	DefaultAddress = "http://localhost:9091/transmission/rpc"
)

// Config used to configure transmission client
type Config struct {
	// Address defaultt http://localhost:9091/transmission/rpc
	Address  string
	User     string
	Password string
	// SkipCheckSSL set to true if you use untrusted certificat default false
	SkipCheckSSL bool
}

// Client transmission client
type Client struct {
	Session    *Session
	httpClient *http.Client
	conf       *Config
	sessionID  string
	endpoint   string
}

type getTorrentArg struct {
	Fields []string `json:"fields,omitempty"`
	Ids    []int    `json:"ids,omitempty"`
}

// AddTorrentArg params for Client.AddTorrent
type AddTorrentArg struct {
	// Cookies string
	// DownloadDir path to download the torrent to
	DownloadDir string `json:"download-dir,omitempty"`
	// Filename filename or URL of the .torrent file
	Filename string `json:"filename,omitempty"`
	// Metainfo base64-encoded .torrent content
	Metainfo string `json:"metainfo,omitempty"`
	// Paused if true add torrent paused default false
	Paused bool `json:"paused,omitempty"`
	// PeerLimit maximum number of peers
	PeerLimit int `json:"peer-limit,omitempty"`
	// BandwidthPriority torrent's bandwidth
	BandwidthPriority int `json:",omitempty"`
	// files-wanted
	// files-unwanted
	// priority-high
	// priority-low
	// priority-normal

}

type removeTorrentArg struct {
	Ids             []int `json:"ids,string"`
	DeleteLocalData bool  `json:"delete-local-data,omitempty"`
}

// Request object for API call
type Request struct {
	Method    string      `json:"method"`
	Arguments interface{} `json:"arguments"`
}

// Response object for API cal response
type Response struct {
	Arguments interface{} `json:"arguments"`
	Result    string      `json:"result"`
}

// Do low level function for interact with transmission only take care
// of authentification and session id
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

	//Body copy for login response
	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))

	//Log request for debug
	log.Print(bytes.NewBuffer(b).String())

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

// GetTorrents return list of torrent
func (c *Client) GetTorrents() ([]*Torrent, error) {
	tReq := &Request{
		Arguments: getTorrentArg{
			Fields: torrentGetFields,
		},
		Method: "torrent-get",
	}

	r := &Response{Arguments: &Torrents{}}

	err := c.request(tReq, r)
	if err != nil {
		return nil, err
	}

	t := r.Arguments.(*Torrents).Torrents
	for i := 0; i < len(t); i++ {
		t[i].Client = c
	}
	return t, nil
}

// Add shortcut for Client.AddTorrent
func (c *Client) Add(filename string) (*Torrent, error) {
	args := AddTorrentArg{
		Filename: filename,
	}
	return c.AddTorrent(args)
}

// AddTorrent add torrent from filename or metadata
// see AddTorrentArg for arguments
func (c *Client) AddTorrent(args AddTorrentArg) (*Torrent, error) {
	tReq := &Request{
		Arguments: args,
		Method:    "torrent-add",
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

// RemoveTorrents remove torrents
func (c *Client) RemoveTorrents(torrents []*Torrent, removeData bool) error {
	ids := make([]int, len(torrents))
	for i := range torrents {
		ids[i] = torrents[i].ID
	}
	tReq := &Request{
		Arguments: removeTorrentArg{
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

// BlocklistUpdate update blocklist and return blocklist rules size
func (c *Client) BlocklistUpdate() (int, error) {
	tReq := &Request{
		Method: "blocklist-update",
	}
	type update struct {
		BlocklistSize int `json:"blocklist-size"`
	}
	r := &Response{Arguments: &update{}}
	err := c.request(tReq, r)
	if err != nil {
		return 0, err
	}
	s := r.Arguments.(*update)
	return s.BlocklistSize, nil
}

// PortTest tests to see if your incoming peer port is accessible
// from the outside world.
func (c *Client) PortTest() (bool, error) {
	tReq := &Request{
		Method: "port-test",
	}
	type rep struct {
		IsOpen bool `json:"port-is-open"`
	}
	r := &Response{Arguments: &rep{}}
	err := c.request(tReq, r)
	if err != nil {
		return false, err
	}
	s := r.Arguments.(*rep)
	return s.IsOpen, nil
}

// FreeSpace tests how much free space is available in a
// client-specified folder.
func (c *Client) FreeSpace(path string) (int, error) {
	type arg struct {
		Path string `json:"path"`
	}
	tReq := &Request{
		Arguments: arg{path},
		Method:    "free-space",
	}
	type rep struct {
		Path      string `json:"path"`
		SizeBytes int    `json:"size-bytes"`
	}
	r := &Response{Arguments: &rep{}}
	err := c.request(tReq, r)
	if err != nil {
		return 0, err
	}
	s := r.Arguments.(*rep)
	return s.SizeBytes, nil
}

func (c *Client) QueueMoveTop(torrents []*Torrent) error {
	return c.queueAction("queue-move-top", torrents)
}
func (c *Client) QueueMoveUp(torrents []*Torrent) error {
	return c.queueAction("queue-move-up", torrents)
}
func (c *Client) QueueMoveDown(torrents []*Torrent) error {
	return c.queueAction("queue-move-down", torrents)
}
func (c *Client) QueueMoveBottom(torrents []*Torrent) error {
	return c.queueAction("queue-move-bottom", torrents)
}

func (c *Client) queueAction(method string, torrents []*Torrent) error {
	ids := make([]int, len(torrents))
	for i := range torrents {
		ids[i] = torrents[i].ID
	}
	type arg struct {
		Ids []int `json:"ids"`
	}
	tReq := &Request{
		Arguments: arg{ids},
		Method:    method,
	}
	r := &Response{}
	err := c.request(tReq, r)
	if err != nil {
		return err
	}
	return nil
}

// New create a new transmission client
func New(conf Config) (*Client, error) {
	httpClient := &http.Client{}
	if conf.SkipCheckSSL {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		httpClient = &http.Client{Transport: tr}
	}
	if conf.Address == "" {
		conf.Address = DefaultAddress
	}
	client := Client{
		conf:       &conf,
		httpClient: httpClient,
		endpoint:   conf.Address,
		Session:    &Session{},
	}
	client.Session.Client = &client
	return &client, nil
}

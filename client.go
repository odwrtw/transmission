package transmission

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	// ErrDuplicateTorrent returned when the torrent is already added
	ErrDuplicateTorrent = errors.New("Torrent already added")
)

// Config used to configure transmission client
type Config struct {
	Address  string
	User     string
	Password string
	// HTTPClient specify your own if you need default: http.Client
	HTTPClient *http.Client
}

// Client transmission client
type Client struct {
	Session   *Session
	sessionID string
	*Config
}

// AddTorrentArg params for Client.AddTorrent
type AddTorrentArg struct {
	//The format of the "cookies" should be NAME=CONTENTS, where NAME is the
	//cookie name and CONTENTS is what the cookie should contain.  Set multiple
	//cookies like this: "name1=content1; name2=content2;" etc.
	//<http://curl.haxx.se/libcurl/c/curl_easy_setopt.html#CURLOPTCOOKIE>
	Cookies string `json:"cookies,omitempty"`
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
	BandwidthPriority int   `json:",omitempty"`
	FilesWanted       []int `json:"files-wanted,omitempty"`
	FilesUnwanted     []int `json:"files-unwanted,omitempty"`
	PriorityHigh      []int `json:"priority-high,omitempty"`
	PriorityLow       []int `json:"priority-low,omitempty"`
	PriorityNormal    []int `json:"priority-normal,omitempty"`
}

// Request object for API call
type Request struct {
	Method    string      `json:"method"`
	Arguments interface{} `json:"arguments"`
}

// Response object for API call response
type Response struct {
	Arguments interface{} `json:"arguments"`
	Result    string      `json:"result"`
}

// Do low level function for interact with transmission only take care of
// authentification and session id
func (c *Client) Do(req *http.Request, retry bool) (*http.Response, error) {
	if c.User != "" && c.Password != "" {
		req.SetBasicAuth(c.User, c.Password)
	}
	if c.sessionID != "" {
		req.Header.Add("X-Transmission-Session-Id", c.sessionID)
	}

	//Body copy for replay it if needed
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	req.Body.Close()
	req.Body = ioutil.NopCloser(bytes.NewBuffer(b))

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Most Transmission RPC servers require a X-Transmission-Session-Id header
	// to be sent with requests, to prevent CSRF attacks.  When your request
	// has the wrong id -- such as when you send your first request, or when
	// the server expires the CSRF token -- the Transmission RPC server will
	// return an HTTP 409 error with the right X-Transmission-Session-Id in its
	// own headers.  So, the correct way to handle a 409 response is to update
	// your X-Transmission-Session-Id and to resend the previous request.
	if resp.StatusCode == http.StatusConflict && retry {
		c.sessionID = resp.Header.Get("X-Transmission-Session-Id")
		req.Body = ioutil.NopCloser(bytes.NewBuffer(b))
		return c.Do(req, false)
	}

	return resp, nil
}

func (c *Client) request(tReq *Request, tResp *Response) error {

	data, err := json.Marshal(tReq)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", c.Address, bytes.NewReader(data))
	if err != nil {
		return err
	}

	resp, err := c.Do(req, true)
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
	type arg struct {
		Fields []string `json:"fields,omitempty"`
		Ids    []int    `json:"ids,omitempty"`
	}

	tReq := &Request{
		Arguments: arg{
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
	for i := range t {
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

// AddTorrent add torrent from filename or metadata see AddTorrentArg for
// arguments
func (c *Client) AddTorrent(args AddTorrentArg) (*Torrent, error) {
	tReq := &Request{
		Arguments: args,
		Method:    "torrent-add",
	}
	// TODO: When there is an error like that, the name of the error is in the
	// name of the key of the JSON
	// Here we only get the torrent-added response, but in other cases the
	// response could be for example torrent-duplicate
	// Need to unmarshal with json.RawMessage
	type added struct {
		Torrent *Torrent `json:"torrent-added"`
	}
	r := &Response{Arguments: &added{}}
	err := c.request(tReq, r)
	if err != nil {
		return nil, err
	}
	t := r.Arguments.(*added)

	// If it's a success but we didn't add any torrent, it's because the
	// torrent is already added
	if t.Torrent == nil {
		return nil, ErrDuplicateTorrent
	}

	t.Torrent.Client = c
	return t.Torrent, nil
}

// RemoveTorrents remove torrents
func (c *Client) RemoveTorrents(torrents []*Torrent, removeData bool) error {
	ids := make([]int, len(torrents))
	for i := range torrents {
		ids[i] = torrents[i].ID
	}

	type arg struct {
		Ids             []int `json:"ids,string"`
		DeleteLocalData bool  `json:"delete-local-data,omitempty"`
	}

	tReq := &Request{
		Arguments: arg{
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

// PortTest tests to see if your incoming peer port is accessible from the
// outside world.
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

// FreeSpace tests how much free space is available in a client-specified
// folder.
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

// QueueMoveTop moves torrents to top of the queue
func (c *Client) QueueMoveTop(torrents []*Torrent) error {
	return c.queueAction("queue-move-top", torrents)
}

// QueueMoveUp moves torrents up in the queue
func (c *Client) QueueMoveUp(torrents []*Torrent) error {
	return c.queueAction("queue-move-up", torrents)
}

// QueueMoveDown moves torrents down in the queue
func (c *Client) QueueMoveDown(torrents []*Torrent) error {
	return c.queueAction("queue-move-down", torrents)
}

// QueueMoveBottom moves torrents to botton of the queue
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
	if conf.HTTPClient == nil {
		conf.HTTPClient = &http.Client{}
	}
	client := Client{
		Session: &Session{},
		Config:  &conf,
	}
	client.Session.Client = &client
	return &client, nil
}

package transmission

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"sync"
)

var tokenRegexp *regexp.Regexp

var (
	// ErrDuplicateTorrent is the error returned when trying to add a torrent
	// already in transmission
	ErrDuplicateTorrent = errors.New("torrent already added")
)

const (
	// StatuStopped status when torrent is stopped
	StatuStopped = 0
	// StatusCheckWait status when torrent is queued to while checking files
	StatusCheckWait = 1
	// StatusCheck status when torrent fies are being checked
	StatusCheck = 2
	// StatusDownloadWait status when torrent queued to be downloaded
	StatusDownloadWait = 3
	// StatusDownload status when torrent is downloading
	StatusDownload = 4
	// StatusSeedWait status when torrent is in queue to seed
	StatusSeedWait = 5
	// StatusSeed status when torrent is seeding
	StatusSeed = 6
)

func init() {
	// Regexp to get the header token
	tokenRegexp = regexp.MustCompile("X-Transmission-Session-Id:\\s([^<]+)")
}

// Transmission type
type Transmission struct {
	Endpoint     string
	Username     string
	Password     string
	Token        string
	SkipCheckSSL bool

	once       *sync.Once
	tokenError error
}

// PostData represents the data to post
type PostData struct {
	Arguments PostArguments `json:"arguments"`
	Method    string        `json:"method"`
}

// PostArguments represents the post arguments
type PostArguments struct {
	Fields   []string `json:"fields,omitempty"`
	Filename string   `json:"filename,omitempty"`
	Ids      []int    `json:"ids,omitempty"`
}

// Result represents the result from the rpc call
type Result struct {
	Arguments ResultArguments `json:"arguments"`
	Status    string          `json:"result"`
}

// ResultArguments represents the result arguments form the result
type ResultArguments struct {
	Torrents     []*ResultTorrent `json:"torrents"`
	TorrentAdded *ResultTorrent   `json:"torrent-added"`
}

// ResultTorrent represents a torrent form the result
type ResultTorrent struct {
	Hash        string  `json:"hashString"`
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	RatioDone   float64 `json:"percentDone"`
	Status      int     `json:"status"`
	UploadRatio float64 `json:"uploadRatio"`
}

func (t *Transmission) getToken() {
	// Make the get request to get the token
	body, err := t.makeGetRequest()

	token := tokenRegexp.FindSubmatch(body)
	if token == nil {
		t.tokenError = err
		return
	}

	t.Token = string(token[1])
}

// makeHTTPRequest will make a HTTP request on the endpoint
func (t *Transmission) makeHTTPRequest(httpType string, data []byte) ([]byte, error) {
	// Create Transport
	tr := &http.Transport{}
	if t.SkipCheckSSL {
		tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	// Create Client
	client := &http.Client{Transport: tr}

	// Create Request
	req, err := http.NewRequest(httpType, t.Endpoint, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	// Add Header
	req.Header.Add("Content-Type", "application/json")

	// Add token if exists
	if t.Token != "" {
		req.Header.Add("X-Transmission-Session-Id", t.Token)
	}

	// Add auth if present
	if t.Password != "" {
		req.SetBasicAuth(t.Username, t.Password)
	}

	// Make the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Return the response
	return body, nil
}

// makeGetRequest will make a simple basic GET request on the endpoint
func (t *Transmission) makeGetRequest() ([]byte, error) {
	return t.makeHTTPRequest("GET", nil)
}

// Post send post data to the rpc interface
func (t *Transmission) Post(postData *PostData) (*Result, error) {
	// Get communication token
	t.once.Do(t.getToken)
	if t.tokenError != nil {
		return nil, t.tokenError
	}

	// Encode post data as json
	data, err := json.Marshal(postData)
	if err != nil {
		return nil, err
	}

	// Make the request on the endpoint
	body, err := t.makeHTTPRequest("POST", data)
	if err != nil {
		return nil, err
	}

	// Unmarshal the result
	var result *Result
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	// Check result status
	switch result.Status {
	case "success":
		return result, nil
	case "duplicate torrent":
		return nil, ErrDuplicateTorrent
	default:
		return nil, fmt.Errorf("Failed to post data..\n%#v\n", result)
	}
}

// GetList gets the list of the current torrents
func (t *Transmission) GetList() ([]*ResultTorrent, error) {
	postData := &PostData{
		Arguments: PostArguments{
			Fields: []string{
				"hashString",
				"id",
				"name",
				"percentDone",
				"status",
				"uploadRatio",
			},
		},
		Method: "torrent-get",
	}

	result, err := t.Post(postData)
	if err != nil {
		return nil, err
	}

	return result.Arguments.Torrents, err
}

// AddTorrent add a new torrent to transmission from a magnet or a .torrent URL
func (t *Transmission) AddTorrent(filename string) (*ResultTorrent, error) {
	postData := &PostData{
		Arguments: PostArguments{
			Filename: filename,
		},
		Method: "torrent-add",
	}

	result, err := t.Post(postData)
	if err != nil {
		return nil, err
	}

	return result.Arguments.TorrentAdded, nil
}

// RemoveTorrents remove all the torrents with the given ids
func (t *Transmission) RemoveTorrents(ids []int) error {
	postData := &PostData{
		Arguments: PostArguments{
			Ids: ids,
		},
		Method: "torrent-remove",
	}

	_, err := t.Post(postData)
	if err != nil {
		return err
	}

	return nil
}

// New return a new pointer of transmission
func New(endpoint string) *Transmission {
	return &Transmission{
		Endpoint:     endpoint,
		SkipCheckSSL: false,
		once:         &sync.Once{},
	}
}

// NewWithAuth returns a new pointer of transmission with Auth
func NewWithAuth(endpoint string, username string, password string) *Transmission {
	return &Transmission{
		Endpoint:     endpoint,
		Password:     password,
		SkipCheckSSL: false,
		Username:     username,
		once:         &sync.Once{},
	}
}

package transmission

const (
	// Status of torrents
	StatusStopped         = 0
	StatusCheckPending    = 1
	StatusChecking        = 2
	StatusDownloadPending = 3
	StatusDownloading     = 4
	StatusSeedPending     = 5
	Statusseeding         = 6
)

// Torrents a lis of Torrents
type Torrents struct {
	Torrents []*Torrent `json:"torrents"`
}

// Torrent represent a torrent present in transmission
type Torrent struct {
	Client                  *Client `json:"-"`
	ActivityDate            int
	AddedDate               int
	BandwidthPriority       int
	Comment                 string
	CorruptEver             int
	Creator                 string
	DateCreated             int
	DesiredAvailable        int
	DoneDate                int
	DownloadDir             string
	DownloadedEver          int
	DownloadLimit           int
	DownloadLimited         bool
	Error                   int
	ErrorString             string
	Eta                     int
	EtaIdle                 int
	Files                   *[]File
	FileStats               *[]FileStats
	HashString              string
	HaveUnchecked           int
	HaveValid               int
	HonorsSessionLimits     bool
	ID                      int
	IsFinished              bool
	IsPrivate               bool
	IsStalled               bool
	LeftUntilDone           int
	MagnetLink              string
	ManualAnnounceTime      int
	MaxConnectedPeers       int
	MetadataPercentComplete float64
	Name                    string
	Peerlimit               int
	Peers                   *[]Peers
	PeersConnected          int
	PeersFrom               PeersFrom
	PeersGettingFromUs      int
	PeersSendingToUs        int
	PercentDone             float64
	Pieces                  string
	PieceCount              int
	PieceSize               int
	// Priorities         *[]Priorities
	QueuePosition      int
	RateDownload       int
	RateUpload         int
	RecheckProgress    float64
	SecondsDownloading int
	SecondsSeeding     int
	SeedIdleLimit      int
	SeedIdleMode       int
	SeedRatioLimit     float64
	SeedRatioMode      int
	SizeWhenDone       int
	StartDate          int
	Status             int
	Trackers           *[]Trackers
	// TrackerStats  *[]TrackerStats
	TotalSize     int
	TorrentFile   string
	UploadedEver  int
	UploadLimit   int
	UploadLimited bool
	UploadRatio   float64
	// wanted                       array
	webseeds            interface{}
	WebseedsSendingToUs int
}

func (t *Torrent) torrentAction(method string) error {
	type Arg struct {
		Ids int `json:"ids"`
	}
	tReq := &Request{
		Arguments: Arg{
			Ids: t.ID,
		},
		Method: method,
	}

	r := &Response{}
	err := t.Client.request(tReq, r)
	if err != nil {
		return err
	}
	return nil

}

// Start torrent
func (t *Torrent) Start() error {
	return t.torrentAction("torrent-start")
}

// StartNow torrent
func (t *Torrent) StartNow() error {
	return t.torrentAction("torrent-start-now")
}

// Stop torrent
func (t *Torrent) Stop() error {
	return t.torrentAction("torrent-stop")
}

// Verify torrent
func (t *Torrent) Verify() error {
	return t.torrentAction("torrent-verify")
}

// Reannounce torrent
func (t *Torrent) Reannounce() error {
	return t.torrentAction("torrent-reannounce")
}

// PathRename renames a file or directory in a torrent.
func (t *Torrent) PathRename(path string, newPath string) error {
	type arg struct {
		Ids  []int  `json:"ids,string"`
		Path string `json:"path"`
		Name string `json:"name"`
	}
	tReq := &Request{
		Arguments: arg{
			Ids:  []int{t.ID},
			Path: path,
			Name: newPath,
		},
		Method: "torrent-rename-path",
	}

	r := &Response{}
	err := t.Client.request(tReq, r)
	if err != nil {
		return err
	}
	return nil
}

// Update torrent information from transmission
func (t *Torrent) Update() error {
	type Arg struct {
		Ids    int      `json:"ids"`
		Fields []string `json:"fields,omitempty"`
	}
	tReq := &Request{
		Arguments: Arg{
			Ids:    t.ID,
			Fields: torrentGetFields,
		},
		Method: "torrent-get",
	}
	r := &Response{
		Arguments: &Torrents{
			Torrents: []*Torrent{t},
		},
	}
	err := t.Client.request(tReq, r)
	if err != nil {
		return err
	}
	return nil
}

package transmission

const (
	// StatusStopped stopped
	StatusStopped = 0
	// StatusCheckPending check pending
	StatusCheckPending = 1
	// StatusChecking checking
	StatusChecking = 2
	// StatusDownloadPending download pending
	StatusDownloadPending = 3
	// StatusDownloading downloading
	StatusDownloading = 4
	// StatusSeedPending seed pending
	StatusSeedPending = 5
	// StatusSeeding seeding
	StatusSeeding = 6
)

var torrentGetFields = []string{
	"activityDate",
	"addedDate",
	"bandwidthPriority",
	"comment",
	"corruptEver",
	"creator",
	"dateCreated",
	"desiredAvailable",
	"doneDate",
	"downloadDir",
	"downloadedEver",
	"downloadLimit",
	"downloadLimited",
	"error",
	"errorString",
	"eta",
	"etaIdle",
	"files",
	"fileStats",
	"hashString",
	"haveUnchecked",
	"haveValid",
	"honorsSessionLimits",
	"id",
	"isFinished",
	"isPrivate",
	"isStalled",
	"leftUntilDone",
	"magnetLink",
	"manualAnnounceTime",
	"maxConnectedPeers",
	"metadataPercentComplete",
	"name",
	"peer",
	"peers",
	"peersConnected",
	"peersFrom",
	"peersGettingFromUs",
	"peersSendingToUs",
	"percentDone",
	"pieces",
	"pieceCount",
	"pieceSize",
	"priorities",
	"queuePosition",
	"rateDownload",
	"rateUpload",
	"recheckProgress",
	"secondsDownloading",
	"secondsSeeding",
	"seedIdleLimit",
	"seedIdleMode",
	"seedRatioLimit",
	"seedRatioMode",
	"sizeWhenDone",
	"startDate",
	"status",
	"trackers",
	"trackerStats",
	"totalSize",
	"torrentFile",
	"uploadedEver",
	"uploadLimit",
	"uploadLimited",
	"uploadRatio",
	"wanted",
	"webseeds",
	"webseedsSendingToUs",
}

// Torrents a list of Torrents
type Torrents struct {
	Torrents []*Torrent `json:"torrents"`
}

// SetTorrentArg arguments for Torrent.Set method
type SetTorrentArg struct {
	BandwidthPriority   int      `json:"bandwidthPriority,omitempty"`
	DownloadLimit       int      `json:"downloadLimit,omitempty"`
	DownloadLimited     bool     `json:"downloadLimited,omitempty"`
	FilesWanted         []int    `json:"files-wanted,omitempty"`
	FilesUnwanted       []int    `json:"files-unwanted,omitempty"`
	HonorsSessionLimits bool     `json:"honorsSessionLimits,omitempty"`
	Ids                 int      `json:"ids"`
	Location            string   `json:"location,omitempty"`
	PeerLimit           int      `json:"peer-limit,omitempty"`
	PriorityHigh        []int    `json:"priority-high,omitempty"`
	PriorityLow         []int    `json:"priority-low,omitempty"`
	PriorityNormal      []int    `json:"priority-normal,omitempty"`
	QueuePosition       int      `json:"queuePosition,omitempty"`
	SeedIdleLimit       int      `json:"seedIdleLimit,omitempty"`
	SeedIdleMode        int      `json:"seedIdleMode,omitempty"`
	SeedRatioLimit      float64  `json:"seedRatioLimit,omitempty"`
	SeedRatioMode       int      `json:"seedRatioMode,omitempty"`
	TrackerAdd          []string `json:"trackerAdd,omitempty"`
	TrackerRemove       []int    `json:"trackerRemove,omitempty"`
	// TrackerReplace       `json:"trackerReplace,omitempty"`
	UploadLimit   int  `json:"uploadLimit,omitempty"`
	UploadLimited bool `json:"uploadLimited,omitempty"`
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
	Priorities              []int
	QueuePosition           int
	RateDownload            int
	RateUpload              int
	RecheckProgress         float64
	SecondsDownloading      int
	SecondsSeeding          int
	SeedIdleLimit           int
	SeedIdleMode            int
	SeedRatioLimit          float64
	SeedRatioMode           int
	SizeWhenDone            int
	StartDate               int
	Status                  int
	Trackers                *[]Trackers
	TrackerStats            *[]TrackerStats
	TotalSize               int
	TorrentFile             string
	UploadedEver            int
	UploadLimit             int
	UploadLimited           bool
	UploadRatio             float64
	Wanted                  []int
	Webseeds                []string
	WebseedsSendingToUs     int
}

// File transmission API response
type File struct {
	BytesCompleted int
	Length         int
	Name           string
}

// FileStats transmission API response
type FileStats struct {
	BytesCompleted int
	Wanted         bool
	Priority       int
}

// Peers transmission API response
type Peers struct {
	Address            string
	ClientName         string
	ClientIsChoked     bool
	ClientIsInterested bool
	FlagStr            string
	IsDownloadingFrom  bool
	IsEncrypted        bool
	IsIncoming         bool
	IsUploadingTo      bool
	IsUTP              bool
	PeerIsChoked       bool
	PeerIsInterested   bool
	Port               int
	Progress           float64
	RateToClient       int
	RateToPeer         int
}

// PeersFrom transmission API response
type PeersFrom struct {
	FromCache    int
	FromDht      int
	FromIncoming int
	FromLpd      int
	FromLtep     int
	FromPex      int
	FromTracker  int
}

// TrackerStats transmission API response
type TrackerStats struct {
	Announce              string
	AnnounceState         int
	DownloadCount         int
	HasAnnounced          bool
	HasScraped            bool
	Host                  string
	ID                    int
	IsBackup              bool
	LastAnnouncePeerCount int
	LastAnnounceResult    string
	LastAnnounceStartTime int
	LastAnnounceSucceeded bool
	LastAnnounceTime      int
	LastAnnounceTimedOut  bool
	LastScrapeResult      string
	LastScrapeStartTime   int
	LastScrapeSucceeded   bool
	LastScrapeTime        int
	LastScrapeTimedOut    int
	LeecherCount          int
	NextAnnounceTime      int
	NextScrapeTim         int
	Scrap                 string
	ScrapeState           int
	SeederCount           int
	Tier                  int
}

// Trackers from transmission API response
type Trackers struct {
	Announce string
	ID       int
	Scrape   string
	Tier     int
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

// SetLocation moves a Torrent
// move if true, move from previous location.
// otherwise, search "location" for files
func (t *Torrent) SetLocation(path string, move bool) error {
	type arg struct {
		Ids      []int  `json:"ids,string"`
		Location string `json:"location"`
		Move     bool   `json:"move,omitempty"`
	}
	tReq := &Request{
		Arguments: arg{
			Ids:      []int{t.ID},
			Location: path,
			Move:     move,
		},
		Method: "torrent-set-location",
	}

	r := &Response{}
	err := t.Client.request(tReq, r)
	if err != nil {
		return err
	}
	return nil
}

// Set changes torrent param see SetTorrentArg
func (t *Torrent) Set(arg SetTorrentArg) error {
	arg.Ids = t.ID
	tReq := &Request{
		Arguments: arg,
		Method:    "torrent-set",
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

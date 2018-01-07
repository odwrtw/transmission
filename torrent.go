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

// TorrentMap is a map of Torrents indexed by torrent hash.
type TorrentMap map[string]*Torrent

// SetTorrentArg arguments for Torrent.Set method
type SetTorrentArg struct {
	BandwidthPriority   int64      `json:"bandwidthPriority,omitempty"`
	DownloadLimit       int64      `json:"downloadLimit,omitempty"`
	DownloadLimited     bool     `json:"downloadLimited,omitempty"`
	FilesWanted         []int64    `json:"files-wanted,omitempty"`
	FilesUnwanted       []int64    `json:"files-unwanted,omitempty"`
	HonorsSessionLimits bool     `json:"honorsSessionLimits,omitempty"`
	Ids                 int64      `json:"ids"`
	Location            string   `json:"location,omitempty"`
	PeerLimit           int64      `json:"peer-limit,omitempty"`
	PriorityHigh        []int64    `json:"priority-high,omitempty"`
	PriorityLow         []int64    `json:"priority-low,omitempty"`
	PriorityNormal      []int64    `json:"priority-normal,omitempty"`
	QueuePosition       int64      `json:"queuePosition,omitempty"`
	SeedIdleLimit       int64      `json:"seedIdleLimit,omitempty"`
	SeedIdleMode        int64      `json:"seedIdleMode,omitempty"`
	SeedRatioLimit      float64  `json:"seedRatioLimit,omitempty"`
	SeedRatioMode       int64      `json:"seedRatioMode,omitempty"`
	TrackerAdd          []string `json:"trackerAdd,omitempty"`
	TrackerRemove       []int64    `json:"trackerRemove,omitempty"`
	// TrackerReplace       `json:"trackerReplace,omitempty"`
	UploadLimit   int64  `json:"uploadLimit,omitempty"`
	UploadLimited bool `json:"uploadLimited,omitempty"`
}

// Torrent represent a torrent present in transmission
type Torrent struct {
	Client                  *Client `json:"-"`
	ActivityDate            int64
	AddedDate               int64
	BandwidthPriority       int64
	Comment                 string
	CorruptEver             int64
	Creator                 string
	DateCreated             int64
	DesiredAvailable        int64
	DoneDate                int64
	DownloadDir             string
	DownloadedEver          int64
	DownloadLimit           int64
	DownloadLimited         bool
	Error                   int64
	ErrorString             string
	Eta                     int64
	EtaIdle                 int64
	Files                   *[]File
	FileStats               *[]FileStats
	HashString              string
	HaveUnchecked           int64
	HaveValid               int64
	HonorsSessionLimits     bool
	ID                      int64
	IsFinished              bool
	IsPrivate               bool
	IsStalled               bool
	LeftUntilDone           int64
	MagnetLink              string
	ManualAnnounceTime      int64
	MaxConnectedPeers       int64
	MetadataPercentComplete float64
	Name                    string
	Peerlimit               int64
	Peers                   *[]Peers
	PeersConnected          int64
	PeersFrom               PeersFrom
	PeersGettingFromUs      int64
	PeersSendingToUs        int64
	PercentDone             float64
	Pieces                  string
	PieceCount              int64
	PieceSize               int64
	Priorities              []int64
	QueuePosition           int64
	RateDownload            int64
	RateUpload              int64
	RecheckProgress         float64
	SecondsDownloading      int64
	SecondsSeeding          int64
	SeedIdleLimit           int64
	SeedIdleMode            int64
	SeedRatioLimit          float64
	SeedRatioMode           int64
	SizeWhenDone            int64
	StartDate               int64
	Status                  int64
	Trackers                *[]Trackers
	TrackerStats            *[]TrackerStats
	TotalSize               int64
	TorrentFile             string
	UploadedEver            int64
	UploadLimit             int64
	UploadLimited           bool
	UploadRatio             float64
	Wanted                  []int64
	Webseeds                []string
	WebseedsSendingToUs     int64
}

// File transmission API response
type File struct {
	BytesCompleted int64
	Length         int64
	Name           string
}

// FileStats transmission API response
type FileStats struct {
	BytesCompleted int64
	Wanted         bool
	Priority       int64
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
	Port               int64
	Progress           float64
	RateToClient       int64
	RateToPeer         int64
}

// PeersFrom transmission API response
type PeersFrom struct {
	FromCache    int64
	FromDht      int64
	FromIncoming int64
	FromLpd      int64
	FromLtep     int64
	FromPex      int64
	FromTracker  int64
}

// TrackerStats transmission API response
type TrackerStats struct {
	Announce              string
	AnnounceState         int64
	DownloadCount         int64
	HasAnnounced          bool
	HasScraped            bool
	Host                  string
	ID                    int64
	IsBackup              bool
	LastAnnouncePeerCount int64
	LastAnnounceResult    string
	LastAnnounceStartTime int64
	LastAnnounceSucceeded bool
	LastAnnounceTime      int64
	LastAnnounceTimedOut  bool
	LastScrapeResult      string
	LastScrapeStartTime   int64
	LastScrapeSucceeded   bool
	LastScrapeTime        int64
	LastScrapeTimedOut    int64
	LeecherCount          int64
	NextAnnounceTime      int64
	NextScrapeTim         int64
	Scrap                 string
	ScrapeState           int64
	SeederCount           int64
	Tier                  int64
}

// Trackers from transmission API response
type Trackers struct {
	Announce string
	ID       int64
	Scrape   string
	Tier     int64
}

func (t *Torrent) torrentAction(method string) error {
	type Arg struct {
		Ids int64 `json:"ids"`
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
		Ids  []int64  `json:"ids,string"`
		Path string `json:"path"`
		Name string `json:"name"`
	}
	tReq := &Request{
		Arguments: arg{
			Ids:  []int64{t.ID},
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
		Ids      []int64  `json:"ids,string"`
		Location string `json:"location"`
		Move     bool   `json:"move,omitempty"`
	}
	tReq := &Request{
		Arguments: arg{
			Ids:      []int64{t.ID},
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
		Ids    int64      `json:"ids"`
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

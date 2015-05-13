package transmission

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

type File struct {
	BytesCompleted int
	Length         int
	Name           string
}

type FileStats struct {
	BytesCompleted int
	Wanted         bool
	Priority       int
}

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

type PeersFrom struct {
	FromCache    int
	FromDht      int
	FromIncoming int
	FromLpd      int
	FromLtep     int
	FromPex      int
	FromTracker  int
}

type Trackers struct {
	Announce string
	Id       int
	Scrape   string
	Tier     int
}

// type TrackerStats struct {
// 	Announce              string
// 	AnnounceState         int
// 	DownloadCount         int
// 	HasAnnounced          bool
// 	HasScraped            bool
// 	Host                  string
// 	Id                    int
// 	IsBackup              bool
// 	LastAnnouncePeerCount int
// 	LastAnnounceResult    string
// 	LastAnnounceStartTime int
// 	LastAnnounceSucceeded bool
// 	LastAnnounceTime      int
// 	LastAnnounceTimedOut  bool
// 	LastScrapeResult      string
// 	LastScrapeStartTime   int
// 	LastScrapeSucceeded   bool
// 	LastScrapeTime        int
// 	LastScrapeTimedOut    bool
// 	LeecherCount          int
// 	NextAnnounceTime      int
// 	NextScrapeTim         int
// 	Scrap                 string
// 	ScrapeState           int
// 	SeederCount           int
// 	Tier                  int
// }

type Torrents struct {
	Torrents *[]Torrent `json:"torrents"`
}

type Torrent struct {
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
	Id                      int
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

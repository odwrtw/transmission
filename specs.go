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

type Torrents struct {
	Torrents *[]Torrent `json:"torrents"`
}

type Torrent struct {
	ActivityDate      int
	AddedDate         int
	BandwidthPriority int
	Comment           string
	CorruptEver       int
	Creator           string
	DateCreated       int
	DesiredAvailable  int
	DoneDate          int
	DownloadDir       string
	DownloadedEver    int
	DownloadLimit     int
	DownloadLimited   bool
	Error             int
	ErrorString       string
	Eta               int
	EtaIdle           int
	// files                        array
	// fileStats                    array
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
	// peers                        array
	PeersConnected int
	// peersFrom                    object
	PeersGettingFromUs int
	PeersSendingToUs   int
	PercentDone        float64
	Pieces             string
	PieceCount         int
	PieceSize          int
	// priorities                   array
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
	// trackers                     array
	// trackerStats                 array
	TotalSize     int
	TorrentFile   string
	UploadedEver  int
	UploadLimit   int
	UploadLimited bool
	UploadRatio   float64
	// wanted                       array
	// webseeds                     array
	WebseedsSendingToUs int
}

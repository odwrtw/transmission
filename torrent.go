package transmission

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

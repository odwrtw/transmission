package transmission

// Session object contain information about transmission
// session and interact with it
type Session struct {
	Client                    *Client `json:"-"`
	AltSpeedDown              int     `jsonn:"alt-speed-down"`
	AltSpeedEnabled           bool    `json:"alt-speed-enabled"`
	AltSpeedTimeBegin         int     `json:"alt-speed-time-begin"`
	AltSpeedTimeEnabled       bool    `json:"alt-speed-time-enabled"`
	AltSpeedTimeEnd           int     `json:"alt-speed-time-end"`
	AltSpeedTimeDay           int     `json:"alt-speed-time-day"`
	AltSpeedUp                int     `json:"alt-speed-up"`
	BlocklistURL              string  `json:"blocklist-url"`
	BlocklistEnabled          bool    `json:"blocklist-enabled"`
	BlocklistSize             int     `json:"blocklist-size"`
	CacheSizeMb               int     `json:"cache-size-mb"`
	ConfigDir                 string  `json:"config-dir"`
	DownloadDir               string  `json:"download-dir"`
	DownloadQueueSize         int     `json:"download-queue-size"`
	DownloadQueueEnabled      bool    `json:"download-queue-enabled"`
	DhtEnabled                bool    `json:"dht-enabled"`
	Encryption                string  `json:"encryption"`
	IdleSeedingLimit          int     `json:"idle-seeding-limit"`
	IdleSeedingLimitEnabled   bool    `json:"idle-seeding-limit-enabled"`
	IncompleteDir             string  `json:"incomplete-dir"`
	IncompleteDirEnabled      bool    `json:"incomplete-dir-enabled"`
	LpdEnabled                bool    `json:"lpd-enabled"`
	PeerLimitGlobal           int     `json:"peer-limit-global"`
	PeerLimitPerTorrent       int     `json:"peer-limit-per-torrent"`
	PexEnabled                bool    `json:"pex-enabled"`
	PeerPort                  int     `json:"peer-port"`
	PeerPortRandomOnStart     bool    `json:"peer-port-random-on-start"`
	PortForwardingEnabled     bool    `json:"port-forwarding-enabled"`
	QueueStalledEnabled       bool    `json:"queue-stalled-enabled"`
	QueueStalledMinutes       int     `json:"queue-stalled-minutes"`
	RenamePartialFiles        bool    `json:"rename-partial-files"`
	RPCVersion                int     `json:"rpc-version"`
	RPCVersionMinimum         int     `json:"rpc-version-minimum"`
	ScriptTorrentDoneFilename string  `json:"script-torrent-done-filename"`
	ScriptTorrentDoneEnabled  bool    `json:"script-torrent-done-enabled"`
	SeedRatioLimit            float64 `json:"seedRatioLimit"`
	SeedRatioLimited          bool    `json:"seedRatioLimited"`
	SeedQueueSize             int     `json:"seed-queue-size"`
	SeedQueueEnabled          bool    `json:"seed-queue-enabled"`
	SpeedLimitDown            int     `json:"speed-limit-down"`
	SpeedLimitDownEnabled     bool    `json:"speed-limit-down-enabled"`
	SpeedLimitUp              int     `json:"speed-limit-up"`
	SpeedLimitUpEnabled       bool    `json:"speed-limit-up-enabled"`
	StartAddedTorrents        bool    `json:"start-added-torrents"`
	TrashOriginalTorrentFiles bool    `json:"trash-original-torrent-files"`
	// Units                           object `json:"units"`
	UtpEnabled bool   `json:"utp-enabled"`
	Version    string `json:"version"`
}

// Statictics represent session statictics
type Statictics struct {
	ActiveTorrentCount int
	DownloadSpeed      int
	PausedTorrentCount int
	TorrentCount       int
	UploadSpeed        int
	CumulativeStats    *StaticticDetail `json:"cumulative-stats"`
	CurrentStats       *StaticticDetail `json:"current-stats"`
}

// StaticticDetail represent statictics details
type StaticticDetail struct {
	UploadedBytes   int
	DownloadedBytes int
	FilesAdded      int
	SessionCount    int
	SecondsActive   int
}

// Update session information from transmission
func (s *Session) Update() error {
	tReq := &Request{
		Method: "session-get",
	}
	r := &Response{Arguments: s}

	err := s.Client.request(tReq, r)
	if err != nil {
		return err
	}
	return nil
}

// Stats return session statictics
func (s *Session) Stats() (Statictics, error) {
	tReq := &Request{
		Method: "session-stats",
	}

	stat := Statictics{}

	r := &Response{Arguments: &stat}

	err := s.Client.request(tReq, r)
	if err != nil {
		return Statictics{}, err
	}

	return stat, nil
}

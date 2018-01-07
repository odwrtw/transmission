package transmission

// SetSessionArgs arguments for Session.Set
type SetSessionArgs struct {
	AltSpeedDown              int64     `json:"alt-speed-down,omitempty"`
	AltSpeedEnabled           bool    `json:"alt-speed-enabled,omitempty"`
	AltSpeedTimeBegin         int64     `json:"alt-speed-time-begin,omitempty"`
	AltSpeedTimeEnabled       bool    `json:"alt-speed-time-enabled,omitempty"`
	AltSpeedTimeEnd           int64     `json:"alt-speed-time-end,omitempty"`
	AltSpeedTimeDay           int64     `json:"alt-speed-time-day,omitempty"`
	AltSpeedUp                int64     `json:"alt-speed-up,omitempty"`
	BlocklistURL              string  `json:"blocklist-url,omitempty"`
	BlocklistEnabled          bool    `json:"blocklist-enabled,omitempty"`
	CacheSizeMb               int64     `json:"cache-size-mb,omitempty"`
	DownloadDir               string  `json:"download-dir,omitempty"`
	DownloadQueueSize         int64     `json:"download-queue-size,omitempty"`
	DownloadQueueEnabled      bool    `json:"download-queue-enabled,omitempty"`
	DhtEnabled                bool    `json:"dht-enabled,omitempty"`
	Encryption                string  `json:"encryption,omitempty"`
	IdleSeedingLimit          int64     `json:"idle-seeding-limit,omitempty"`
	IdleSeedingLimitEnabled   bool    `json:"idle-seeding-limit-enabled,omitempty"`
	IncompleteDir             string  `json:"incomplete-dir,omitempty"`
	IncompleteDirEnabled      bool    `json:"incomplete-dir-enabled,omitempty"`
	LpdEnabled                bool    `json:"lpd-enabled,omitempty"`
	PeerLimitGlobal           int64     `json:"peer-limit-global,omitempty"`
	PeerLimitPerTorrent       int64     `json:"peer-limit-per-torrent,omitempty"`
	PexEnabled                bool    `json:"pex-enabled,omitempty"`
	PeerPort                  int64     `json:"peer-port,omitempty"`
	PeerPortRandomOnStart     bool    `json:"peer-port-random-on-start,omitempty"`
	PortForwardingEnabled     bool    `json:"port-forwarding-enabled,omitempty"`
	QueueStalledEnabled       bool    `json:"queue-stalled-enabled,omitempty"`
	QueueStalledMinutes       int64     `json:"queue-stalled-minutes,omitempty"`
	RenamePartialFiles        bool    `json:"rename-partial-files,omitempty"`
	ScriptTorrentDoneFilename string  `json:"script-torrent-done-filename,omitempty"`
	ScriptTorrentDoneEnabled  bool    `json:"script-torrent-done-enabled,omitempty"`
	SeedRatioLimit            float64 `json:"seedRatioLimit,omitempty"`
	SeedRatioLimited          bool    `json:"seedRatioLimited,omitempty"`
	SeedQueueSize             int64     `json:"seed-queue-size,omitempty"`
	SeedQueueEnabled          bool    `json:"seed-queue-enabled,omitempty"`
	SpeedLimitDown            int64     `json:"speed-limit-down,omitempty"`
	SpeedLimitDownEnabled     bool    `json:"speed-limit-down-enabled,omitempty"`
	SpeedLimitUp              int64     `json:"speed-limit-up,omitempty"`
	SpeedLimitUpEnabled       bool    `json:"speed-limit-up-enabled,omitempty"`
	StartAddedTorrents        bool    `json:"start-added-torrents,omitempty"`
	TrashOriginalTorrentFiles bool    `json:"trash-original-torrent-files,omitempty"`
	Units                     *Units  `json:"units,omitempty"`
	UtpEnabled                bool    `json:"utp-enabled,omitempty"`
}

// Session object contain information about transmission
// session and interact with it
type Session struct {
	Client                    *Client `json:"-"`
	AltSpeedDown              int64     `jsonn:"alt-speed-down"`
	AltSpeedEnabled           bool    `json:"alt-speed-enabled"`
	AltSpeedTimeBegin         int64     `json:"alt-speed-time-begin"`
	AltSpeedTimeEnabled       bool    `json:"alt-speed-time-enabled"`
	AltSpeedTimeEnd           int64     `json:"alt-speed-time-end"`
	AltSpeedTimeDay           int64     `json:"alt-speed-time-day"`
	AltSpeedUp                int64     `json:"alt-speed-up"`
	BlocklistURL              string  `json:"blocklist-url"`
	BlocklistEnabled          bool    `json:"blocklist-enabled"`
	BlocklistSize             int64     `json:"blocklist-size"`
	CacheSizeMb               int64     `json:"cache-size-mb"`
	ConfigDir                 string  `json:"config-dir"`
	DownloadDir               string  `json:"download-dir"`
	DownloadQueueSize         int64     `json:"download-queue-size"`
	DownloadQueueEnabled      bool    `json:"download-queue-enabled"`
	DhtEnabled                bool    `json:"dht-enabled"`
	Encryption                string  `json:"encryption"`
	IdleSeedingLimit          int64     `json:"idle-seeding-limit"`
	IdleSeedingLimitEnabled   bool    `json:"idle-seeding-limit-enabled"`
	IncompleteDir             string  `json:"incomplete-dir"`
	IncompleteDirEnabled      bool    `json:"incomplete-dir-enabled"`
	LpdEnabled                bool    `json:"lpd-enabled"`
	PeerLimitGlobal           int64     `json:"peer-limit-global"`
	PeerLimitPerTorrent       int64     `json:"peer-limit-per-torrent"`
	PexEnabled                bool    `json:"pex-enabled"`
	PeerPort                  int64     `json:"peer-port"`
	PeerPortRandomOnStart     bool    `json:"peer-port-random-on-start"`
	PortForwardingEnabled     bool    `json:"port-forwarding-enabled"`
	QueueStalledEnabled       bool    `json:"queue-stalled-enabled"`
	QueueStalledMinutes       int64     `json:"queue-stalled-minutes"`
	RenamePartialFiles        bool    `json:"rename-partial-files"`
	RPCVersion                int64     `json:"rpc-version"`
	RPCVersionMinimum         int64     `json:"rpc-version-minimum"`
	ScriptTorrentDoneFilename string  `json:"script-torrent-done-filename"`
	ScriptTorrentDoneEnabled  bool    `json:"script-torrent-done-enabled"`
	SeedRatioLimit            float64 `json:"seedRatioLimit"`
	SeedRatioLimited          bool    `json:"seedRatioLimited"`
	SeedQueueSize             int64     `json:"seed-queue-size"`
	SeedQueueEnabled          bool    `json:"seed-queue-enabled"`
	SpeedLimitDown            int64     `json:"speed-limit-down"`
	SpeedLimitDownEnabled     bool    `json:"speed-limit-down-enabled"`
	SpeedLimitUp              int64     `json:"speed-limit-up"`
	SpeedLimitUpEnabled       bool    `json:"speed-limit-up-enabled"`
	StartAddedTorrents        bool    `json:"start-added-torrents"`
	TrashOriginalTorrentFiles bool    `json:"trash-original-torrent-files"`
	Units                     *Units  `json:"units"`
	UtpEnabled                bool    `json:"utp-enabled"`
	Version                   string  `json:"version"`
}

// Statictics represent session statictics
type Statictics struct {
	ActiveTorrentCount int64
	DownloadSpeed      int64
	PausedTorrentCount int64
	TorrentCount       int64
	UploadSpeed        int64
	CumulativeStats    *StaticticDetail `json:"cumulative-stats"`
	CurrentStats       *StaticticDetail `json:"current-stats"`
}

// StaticticDetail represent statictics details
type StaticticDetail struct {
	UploadedBytes   int64
	DownloadedBytes int64
	FilesAdded      int64
	SessionCount    int64
	SecondsActive   int64
}

// Units in session
type Units struct {
	SpeedUnits  []string `json:"speed-units"`
	SpeedBytes  int64      `json:"speed-bytes"`
	SizeUnits   []string `json:"size-units"`
	SizeBytes   int64      `json:"size-bytes"`
	MemoryUnits []string `json:"memory-units"`
	MemoryBytes int64      `json:"memory-bytes"`
}

// Set set session params see SetSessionArgs
func (s *Session) Set(args SetSessionArgs) error {
	tReq := &Request{
		Arguments: args,
		Method:    "session-set",
	}
	r := &Response{}
	err := s.Client.request(tReq, r)
	if err != nil {
		return err
	}
	return nil
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

// Close tells the transmission session to shut down.
func (s *Session) Close() error {
	tReq := &Request{
		Method: "session-close",
	}
	r := &Response{}

	err := s.Client.request(tReq, r)
	if err != nil {
		return err
	}
	return nil
}

package transmission

// SetSessionArgs arguments for Session.Set
type SetSessionArgs struct {
	AltSpeedDown              int     `json:"alt-speed-down,omitempty"`
	AltSpeedEnabled           bool    `json:"alt-speed-enabled,omitempty"`
	AltSpeedTimeBegin         int     `json:"alt-speed-time-begin,omitempty"`
	AltSpeedTimeEnabled       bool    `json:"alt-speed-time-enabled,omitempty"`
	AltSpeedTimeEnd           int     `json:"alt-speed-time-end,omitempty"`
	AltSpeedTimeDay           int     `json:"alt-speed-time-day,omitempty"`
	AltSpeedUp                int     `json:"alt-speed-up,omitempty"`
	BlocklistURL              string  `json:"blocklist-url,omitempty"`
	BlocklistEnabled          bool    `json:"blocklist-enabled,omitempty"`
	CacheSizeMb               int     `json:"cache-size-mb,omitempty"`
	DownloadDir               string  `json:"download-dir,omitempty"`
	DownloadQueueSize         int     `json:"download-queue-size,omitempty"`
	DownloadQueueEnabled      bool    `json:"download-queue-enabled,omitempty"`
	DhtEnabled                bool    `json:"dht-enabled,omitempty"`
	Encryption                string  `json:"encryption,omitempty"`
	IdleSeedingLimit          int     `json:"idle-seeding-limit,omitempty"`
	IdleSeedingLimitEnabled   bool    `json:"idle-seeding-limit-enabled,omitempty"`
	IncompleteDir             string  `json:"incomplete-dir,omitempty"`
	IncompleteDirEnabled      bool    `json:"incomplete-dir-enabled,omitempty"`
	LpdEnabled                bool    `json:"lpd-enabled,omitempty"`
	PeerLimitGlobal           int     `json:"peer-limit-global,omitempty"`
	PeerLimitPerTorrent       int     `json:"peer-limit-per-torrent,omitempty"`
	PexEnabled                bool    `json:"pex-enabled,omitempty"`
	PeerPort                  int     `json:"peer-port,omitempty"`
	PeerPortRandomOnStart     bool    `json:"peer-port-random-on-start,omitempty"`
	PortForwardingEnabled     bool    `json:"port-forwarding-enabled,omitempty"`
	QueueStalledEnabled       bool    `json:"queue-stalled-enabled,omitempty"`
	QueueStalledMinutes       int     `json:"queue-stalled-minutes,omitempty"`
	RenamePartialFiles        bool    `json:"rename-partial-files,omitempty"`
	ScriptTorrentDoneFilename string  `json:"script-torrent-done-filename,omitempty"`
	ScriptTorrentDoneEnabled  bool    `json:"script-torrent-done-enabled,omitempty"`
	SeedRatioLimit            float64 `json:"seedRatioLimit,omitempty"`
	SeedRatioLimited          bool    `json:"seedRatioLimited,omitempty"`
	SeedQueueSize             int     `json:"seed-queue-size,omitempty"`
	SeedQueueEnabled          bool    `json:"seed-queue-enabled,omitempty"`
	SpeedLimitDown            int     `json:"speed-limit-down,omitempty"`
	SpeedLimitDownEnabled     bool    `json:"speed-limit-down-enabled,omitempty"`
	SpeedLimitUp              int     `json:"speed-limit-up,omitempty"`
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
	Units                     *Units  `json:"units"`
	UtpEnabled                bool    `json:"utp-enabled"`
	Version                   string  `json:"version"`
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

// Units in session
type Units struct {
	SpeedUnits  []string `json:"speed-units"`
	SpeedBytes  int      `json:"speed-bytes"`
	SizeUnits   []string `json:"size-units"`
	SizeBytes   int      `json:"size-bytes"`
	MemoryUnits []string `json:"memory-units"`
	MemoryBytes int      `json:"memory-bytes"`
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

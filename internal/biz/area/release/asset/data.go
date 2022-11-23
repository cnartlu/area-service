package asset

type Asset struct {
	ID            uint64 `json:"id,omitempty"`
	AreaReleaseID uint64 `json:"area_release_id,omitempty"`
	AssetID       uint64 `json:"asset_id,omitempty"`
	AssetName     string `json:"asset_name,omitempty"`
	AssetLabel    string `json:"asset_label,omitempty"`
	AssetState    string `json:"asset_state,omitempty"`
	FilePath      string `json:"file_path,omitempty"`
	FileSize      uint   `json:"file_size,omitempty"`
	DownloadURL   string `json:"download_url,omitempty"`
	Status        Status `json:"status,omitempty"`
	CreateTime    uint64 `json:"create_time,omitempty"`
	UpdateTime    uint64 `json:"update_time,omitempty"`
}

type Status int

const (
	StatusWaitSync   Status = 1
	StatusWaitLoaded Status = 2
	StatusFinished   Status = 99
)

var (
	Status_name = map[Status]string{
		StatusWaitSync:   "待同步",
		StatusWaitLoaded: "待加载",
		StatusFinished:   "已完成",
	}
)

func (s Status) String() string {
	if s, ok := Status_name[s]; ok {
		return s
	}
	return "未知"
}

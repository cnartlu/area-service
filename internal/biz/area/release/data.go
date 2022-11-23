package release

type Release struct {
	ID                 uint64 `json:"id,omitempty"`
	Owner              string `json:"owner,omitempty"`
	Repo               string `json:"repo,omitempty"`
	ReleaseID          uint64 `json:"release_id,omitempty"`
	ReleaseName        string `json:"release_name,omitempty"`
	ReleaseNodeID      string `json:"release_node_id,omitempty"`
	ReleasePublishedAt uint64 `json:"release_published_at,omitempty"`
	ReleaseContent     string `json:"release_content,omitempty"`
	Status             Status `json:"status,omitempty"`
	CreateTime         uint64 `json:"create_time,omitempty"`
	UpdateTime         uint64 `json:"update_time,omitempty"`
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

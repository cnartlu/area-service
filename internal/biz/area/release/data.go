package release

type Release struct {
	// ID
	ID uint64 `json:"id,omitempty"`
	// 创建时间
	CreateTime uint64 `json:"create_time,omitempty"`
	// 更新时间
	UpdateTime uint64 `json:"update_time,omitempty"`
	// 账号
	Owner string `json:"owner,omitempty"`
	// 包名
	Repo string `json:"repo,omitempty"`
	// 发版标识
	ReleaseID uint64 `json:"release_id,omitempty"`
	// 发版名称
	ReleaseName string `json:"release_name,omitempty"`
	// 发版节点
	ReleaseNodeID string `json:"release_node_id,omitempty"`
	// 发版时间
	ReleasePublishedAt uint64 `json:"release_published_at,omitempty"`
	// 发版内容
	ReleaseContent string `json:"release_content,omitempty"`
	// 状态
	Status Status `json:"status,omitempty"`
}

type Status int

const (
	StatusWaitSync   Status = 1
	StatusWaitLoaded Status = 2
	StatusFinished   Status = 99
)

func (s Status) String() string {
	switch s {
	case StatusWaitSync:
		return "待同步"
	case StatusWaitLoaded:
		return "待加载"
	case StatusFinished:
		return "已完成"
	}
	return "未知状态"
}

func (s Status) Values() []Status {
	return []Status{
		StatusWaitSync,
		StatusWaitLoaded,
		StatusFinished,
	}
}

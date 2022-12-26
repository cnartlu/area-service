package splider

import "time"

const (
	SourceGithub = "github.com"
)

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

type Splider struct {
	ID          int
	Source      string
	Owner       string
	Repo        string
	SourceID    uint64
	Title       string
	Draft       bool
	PreRelease  bool
	PublishedAt time.Time
	Status      Status
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

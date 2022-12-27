package asset

import "time"

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

type FindListParams struct {
	CitySpliderID int
	SourceID      uint64
	Status        int
	Page          int
	PageSize      int
}

type Asset struct {
	ID            int
	CitySpliderID int
	SourceID      uint64
	FileTitle     string
	FilePath      string
	FileSize      uint
	Status        Status
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

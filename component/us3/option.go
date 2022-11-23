package us3

import "github.com/cnartlu/area-service/component/filesystem"

var _ filesystem.Option = (*option)(nil)

type option struct {
}

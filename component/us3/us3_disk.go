package us3

import (
	"github.com/cnartlu/area-service/component/filesystem"
	"github.com/mitchellh/mapstructure"
)

func init() {
	filesystem.RegisterTarget(us3Register(0))
}

type us3Register int

func (us3Register) Name() string {
	return "us3"
}

func (l us3Register) Register(data map[string]interface{}) (filesystem.Disk, error) {
	f := Us3{}
	config := &mapstructure.DecoderConfig{
		Metadata:         nil,
		ZeroFields:       true,
		WeaklyTypedInput: true,
		TagName:          "json",
		Result:           &f,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return nil, err
	}
	if err := decoder.Decode(data); err != nil {
		return nil, err
	}
	v := us3Disk{us3: &f}
	return &v, nil
}

var _ filesystem.Disk = (*us3Disk)(nil)

type us3Disk struct {
	us3 *Us3
}

func (u *us3Disk) Exists(filename string, options ...filesystem.HandleFunc) bool {
	return false
}

// Upload 上传文件
func (u *us3Disk) Upload(filename, key string, options ...filesystem.HandleFunc) (filesystem.Result, error) {
	return nil, nil
}

func (u *us3Disk) Url(key string, options ...filesystem.HandleFunc) string {
	return ""
}

func (u *us3Disk) Delete(key string, options ...filesystem.HandleFunc) error {
	return nil
}

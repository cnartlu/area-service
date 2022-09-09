package strings

import (
	"os"
)

// IsDir 是否是目录路径
func IsDir(path string) bool {
	f, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return f.IsDir()
}

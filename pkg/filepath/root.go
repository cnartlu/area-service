package filepath

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var (
	// rootPath 应用程序的根目录
	rootPath string = ""
	// runtimePath 缓存目录
	runtimePath string = ""
)

// RootPath 获取此项目的绝对路径
// 如果是以 go build 生成的二进制文件运行，则返回 二进制文件的目录路径
// 如果是以 go run 运行，则返回在此项目的绝对路径
func RootPath() string {
	if rootPath != "" {
		return rootPath
	}
	var binDir string

	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}

	binDir = filepath.Dir(exePath)

	tmpDir := os.TempDir()
	if strings.Contains(exePath, tmpDir) {
		_, filename, _, ok := runtime.Caller(0)
		if ok {
			binDir = filepath.Dir(filepath.Dir(filepath.Dir(filename)))
		}
	}

	rootPath = binDir

	return binDir
}

// RuntimePath 获取缓存数据的绝对路径
func RuntimePath() string {
	if runtimePath != "" {
		return runtimePath
	}
	runtimePath = filepath.Join(RootPath(), "runtime")
	return runtimePath
}

// RelativePath 返回相对路径
func RelativePath(path string) string {
	rp := RootPath()
	if strings.HasPrefix(path, rp) {
		rpl := len(rp)
		return path[rpl+1:]
	}
	return path
}

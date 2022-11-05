package app

import (
	"os"
	"path/filepath"
	"strings"

	kconfig "github.com/go-kratos/kratos/v2/config"
)

// isFile returns nil if the file exists
func isFile(filename string) error {
	if filename == "" {
		return os.ErrInvalid
	}
	fs, err := os.Stat(filename)
	if err != nil {
		return err
	}
	if fs.IsDir() {
		return os.ErrExist
	}
	return err
}

type Pid struct {
	pid      int
	ppid     int
	filename string
	f        *os.File
}

// Recover 恢复文件
func (p *Pid) Recover() error {
	if strings.HasSuffix(p.filename, ".old") {
		if err := os.Rename(p.filename, strings.TrimSuffix(p.filename, ".old")); err != nil {
			return err
		}
	}
	return nil
}

func (p *Pid) Process() (*os.Process, error) {
	return os.FindProcess(p.pid)
}

func NewPid(appName string, c kconfig.Config) *Pid {
	var pid = Pid{
		pid:      os.Getpid(),
		ppid:     os.Getppid(),
		filename: "",
	}
	if c != nil {
		pidfile, err := c.Value("pid").String()
		if err == nil {
			pid.filename = pidfile
		}
	}
	if pid.filename == "" {
		if appName == "" {
			s := filepath.Base(os.Args[0])
			if i := strings.Index(s, "."); i >= 0 {
				appName = s[:i]
			} else {
				appName = s
			}
		}
		pid.filename = appName + ".pid"
	}
	return &pid
}

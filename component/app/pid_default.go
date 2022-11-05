//go:build !windows

package app

import (
	"os"
	"strconv"
	"syscall"
)

func (p *Pid) Move() error {
	if err := isFile(p.filename); err != nil {
		return err
	}
	var newFilename = p.filename + ".old"
	if err := os.Rename(p.filename, newFilename); err != nil {
		return err
	}
	p.filename = newFilename
	return nil
}

// Remove 删除文件
func (p *Pid) Remove() error {
	var err error
	if p.f != nil {
		err = syscall.Flock(int(p.f.Fd()), syscall.LOCK_UN)
	}
	if err1 := os.Remove(p.filename); err1 != nil && err == nil {
		err = err1
	}
	return err
}

func (p *Pid) WriteFile() error {
	str := strconv.FormatInt(int64(p.pid), 10)
	var err error
	p.f, err = os.OpenFile(p.filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModeExclusive)
	if err != nil {
		return err
	}
	if err := syscall.Flock(int(p.f.Fd()), syscall.LOCK_EX|syscall.LOCK_NB); err != nil {
		if err1 := p.f.Close(); err1 != nil && err == nil {
			err = err1
		}
		if err == syscall.EWOULDBLOCK {
			return err
		}
		return err
	}
	_, err = p.f.WriteString(str)
	if err1 := p.f.Close(); err1 != nil && err == nil {
		err = err1
	}
	return err
}

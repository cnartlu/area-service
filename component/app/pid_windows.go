//go:build windows

package app

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

// Move 移动文件
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
	return os.Remove(p.filename)
}

func (p *Pid) WriteFile() error {
	str := strconv.FormatInt(int64(p.pid), 10)
	f, err := os.OpenFile(p.filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModeExclusive)
	if err != nil {
		return fmt.Errorf("write pid file error: %w", err)
	}
	p.f = f
	fi, _ := filepath.Abs(p.filename)
	fmt.Println("pid", str, "filename", fi)
	_, err = f.WriteString(str)
	if err1 := f.Close(); err1 != nil && err == nil {
		err = err1
	}
	return err
}

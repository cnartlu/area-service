package app

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
)

type pidFile struct {
	filename    string
	oldFilename string
}

func (p *pidFile) GetFilename() string {
	return p.filename
}

func (p *pidFile) GetOldFilename() string {
	return p.oldFilename
}

func (p *pidFile) GetPid() int {
	b, err := os.ReadFile(p.GetFilename())
	if err != nil {
		return 0
	}
	i, err := strconv.ParseInt(string(b), 10, 64)
	if err != nil {
		panic(err)
	}
	return int(i)
}

func (p *pidFile) GetProcess() (*os.Process, error) {
	var pid = p.GetPid()
	if pid != 0 {
		return os.FindProcess(pid)
	}
	return nil, os.NewSyscallError("pid", os.ErrInvalid)
}

func (p *pidFile) WritePid(pid int) error {
	syscall.ForkLock.Lock()
	defer syscall.ForkLock.Unlock()
	dirpath := filepath.Dir(p.GetFilename())
	if f, e := os.Stat(dirpath); e != nil {
		if !os.IsNotExist(e) {
			return e
		}
		if e1 := os.MkdirAll(dirpath, os.ModeDir); e1 != nil {
			return e1
		}
	} else if !f.IsDir() {
		return os.ErrInvalid
	}
	err := os.WriteFile(p.GetFilename(), []byte(strconv.Itoa(pid)), os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (p *pidFile) Remove() error {
	return os.Remove(p.GetFilename())
}

type process struct {
	pid     int
	ppid    int
	process *os.Process
	pidFile *pidFile
}

func (p *process) GetProcess() *os.Process {
	if p != nil {
		return p.process
	}
	return nil
}

func (p *process) GetPidFile() *pidFile {
	if p != nil {
		return p.pidFile
	}
	return nil
}

func (p *process) WritePidFile() error {
	if p != nil {
		return p.pidFile.WritePid(p.pid)
	}
	return nil
}

func (p *process) RemovePidFile() error {
	if p != nil {
		return p.pidFile.Remove()
	}
	return nil
}

func newProcess(a *App) *process {
	var p = process{
		pid:  os.Getpid(),
		ppid: os.Getppid(),
	}
	p.process, _ = os.FindProcess(p.pid)
	var name string
	s := filepath.Base(os.Args[0])
	if i := strings.Index(s, "."); i >= 0 {
		name = s[:i]
	} else {
		name = s
	}
	pf := pidFile{
		filename: filepath.Join(a.GetRuntimePath(), name+".pid"),
	}
	p.pidFile = &pf
	return &p
}

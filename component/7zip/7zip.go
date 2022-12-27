package zip7

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
)

var defaultZip7 = &Zip7{}

type Zip7 struct {
	ctx context.Context
	bin string
}

func (z *Zip7) clone() *Zip7 {
	z1 := *z
	return &z1
}

func (z *Zip7) binFullName() string {
	if z.bin == "" {
		name := BIN_NAME
		z.bin = name
		binname := filepath.Join(filepath.Dir(os.Args[0]), name)
		if f, err1 := os.Stat(binname); err1 == nil {
			if !f.IsDir() {
				z.bin = binname
			}
		}
	}
	return z.bin
}

func WithContext(ctx context.Context) *Zip7 {
	return defaultZip7.WithContext(ctx)
}
func (z *Zip7) WithContext(ctx context.Context) *Zip7 {
	z1 := z.clone()
	z1.ctx = ctx
	return z1
}

func Context() context.Context {
	return defaultZip7.Context()
}
func (z *Zip7) Context() context.Context {
	if z.ctx == nil {
		return context.Background()
	}
	return z.ctx
}

func Extract(filename, path string) error {
	return defaultZip7.Extract(filename, path)
}
func (z *Zip7) Extract(filename, path string) error {
	fullpath := path
	if !filepath.IsAbs(path) {
		if p, err1 := filepath.Abs(path); err1 == nil {
			fullpath = p
		}
	}
	cmd := exec.CommandContext(z.Context(), z.binFullName(), "x", filename, "-y", "-o"+fullpath)
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	return cmd.Run()
}

func New() *Zip7 {
	return &Zip7{}
}

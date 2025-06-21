package pathname

import (
	"os"
	"path/filepath"
	"slices"
)

type Pathname struct {
	path string
}

func NewPathname(path string) Pathname {
	return Pathname{
		path: path,
	}
}

func NewPathnamePtr(path string) *Pathname {
	pathname := NewPathname(path)
	return &pathname
}

func (p *Pathname) ToPath() string {
	if p == nil {
		return ""
	}
	return p.path
}

func (p *Pathname) IsAbsolute() bool {
	if p == nil || p.path == "" {
		return false
	}
	return p.path[0] == '/'
}

func (p *Pathname) IsRelative() bool {
	if p == nil || p.path == "" {
		return false
	}
	return p.path[0] != '/'
}

func (p *Pathname) IsRoot() bool {
	if p == nil || p.path == "" {
		return false
	}
	return p.path == "/"
}

func (p *Pathname) IsExist() bool {
	if p == nil {
		return false
	}

	_, err := os.Stat(p.path)
	return err == nil
}

func (p *Pathname) IsFile() bool {
	if p == nil {
		return false
	}

	info, err := os.Stat(p.path)
	if err != nil {
		return false
	}

	return !info.IsDir()
}

func (p *Pathname) IsDirectory() bool {
	if p == nil {
		return false
	}

	info, err := os.Stat(p.path)
	if err != nil {
		return false
	}

	return info.IsDir()
}

func (p *Pathname) Basename() *Pathname {
	if p == nil {
		return nil
	}

	path := filepath.Base(p.path)
	return NewPathnamePtr(path)
}

func (p *Pathname) Dirname() *Pathname {
	if p == nil {
		return nil
	}

	path := filepath.Dir(p.path)
	return NewPathnamePtr(path)
}

func (p *Pathname) Join(parts ...string) *Pathname {
	if p == nil {
		return nil
	}

	parts = slices.Concat([]string{p.path}, parts)
	path := filepath.Join(parts...)
	return NewPathnamePtr(path)
}

func (p *Pathname) OpenFile(flag int, perm os.FileMode) (*os.File, error) {
	return os.OpenFile(p.ToPath(), flag, perm)
}

func (p *Pathname) Size() (int, error) {
	info, err := os.Stat(p.ToPath())
	if err != nil {
		return 0, err
	}

	return int(info.Size()), nil
}

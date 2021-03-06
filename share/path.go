package share

import (
	"path/filepath"
)

type path struct {
	Root    string
	Config  string
	Runtime string
	Bin     string
}

var Path *path

func InitPath() {
	Path = &path{}
	root, err := filepath.Abs("./../")
	if err != nil {
		panic("path error")
	}
	Path.Root = root
	Path.Bin = root + "\\bin"
	Path.Config = root + "\\config"
	Path.Runtime = root + "\\runtime"
}

package FileSystem

import (
	"github.com/RENCI/GoUtils/FileSystem/Path"
	"os"
)

type FileInfo struct {
	Path string
}

func (this *FileInfo) Create() error {
	f, err := os.Create(this.Path)
	defer f.Close()
	return err
}

func (this *FileInfo) Delete() error {
	err := os.Remove(this.Path)
	return err
}

func (this *FileInfo) Exists() (bool, error) {
	return Path.Exists(this.Path)
}

func (this *FileInfo) Name() string {
	return Path.GetFileName(this.Path)
}

func (this *FileInfo) Directory() DirectoryInfo {
	res := DirectoryInfo{
		Path: Path.GetDirectoryParent(this.Path),
	}

	return res
}

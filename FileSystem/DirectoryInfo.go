package FileSystem

import (
	"github.com/RENCI/GoUtils/Collections/List"
	"github.com/RENCI/GoUtils/FileSystem/Path"
	"path/filepath"
)

type DirectoryInfo struct {
	Path string
}

func (this *DirectoryInfo) GetDirectories() (List.List[DirectoryInfo], error) {
	dirs, err := GetDirectories(this.Path)
	res := List.NewList[DirectoryInfo](dirs.Size())

	for i := 0; i < dirs.Size(); i++ {
		res.Set(i, DirectoryInfo{Path: dirs.Get(i)})
	}
	return res, err
}

func (this *DirectoryInfo) GetParent() DirectoryInfo {
	parentPath := Path.GetDirectoryParent(this.Path)
	di := DirectoryInfo{Path: parentPath}
	return di
}

func (this *DirectoryInfo) CreateSubdirectory(path string) (DirectoryInfo, error) {
	abspath := Path.Combine(this.Path, path)
	err := CreateSubdirectiry(abspath)
	di := DirectoryInfo{Path: abspath}
	return di, err
}

func (this *DirectoryInfo) Create() error {
	err := Create(this.Path)
	return err
}

func (this *DirectoryInfo) Delete() error {
	err := Delete(this.Path)
	return err
}

func (this *DirectoryInfo) GetAbsPath() (string, error) {
	return filepath.Abs(this.Path)
}

func (this *DirectoryInfo) GetFiles() (List.List[FileInfo], error) {
	dirs, err := GetFiles(this.Path)
	res := List.NewList[FileInfo](dirs.Size())

	for i := 0; i < dirs.Size(); i++ {
		res.Set(i, FileInfo{Path: dirs.Get(i)})
	}
	return res, err
}

func (this *DirectoryInfo) GetName() {

}

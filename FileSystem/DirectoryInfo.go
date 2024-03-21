package FileSystem

import (
	"github.com/RENCI/GoUtils/Collections/List"
	"path/filepath"
)

type DirectoryInfo struct {
	Path       string
	_directory DirectoryObj
	_path      PathObj
}

func DirectoryInfo_New(path string) *DirectoryInfo {
	di := DirectoryInfo{Path: path}
	return &di
}

func (this *DirectoryInfo) GetDirectories() (List.List[DirectoryInfo], error) {
	dirs, err := this._directory.GetDirectories(this.Path)
	res := List.NewList[DirectoryInfo](dirs.Size())

	for i := 0; i < dirs.Size(); i++ {
		res.Set(i, DirectoryInfo{Path: dirs.Get(i)})
	}
	return res, err
}

func (this *DirectoryInfo) GetParent() DirectoryInfo {
	parentPath := this._path.GetDirectoryParent(this.Path)
	di := DirectoryInfo{Path: parentPath}
	return di
}

func (this *DirectoryInfo) CreateSubdirectory(path string) (DirectoryInfo, error) {
	abspath := this._path.Combine(this.Path, path)
	err := this._directory.CreateSubdirectiry(abspath)
	di := DirectoryInfo{Path: abspath}
	return di, err
}

func (this *DirectoryInfo) Create() error {
	err := this._directory.Create(this.Path)
	return err
}

func (this *DirectoryInfo) Delete() error {
	err := this._directory.Delete(this.Path)
	return err
}

func (this *DirectoryInfo) GetAbsPath() (string, error) {
	return filepath.Abs(this.Path)
}

func (this *DirectoryInfo) GetFiles() (List.List[FileInfo], error) {
	dirs, err := this._directory.GetFiles(this.Path)
	res := List.NewList[FileInfo](dirs.Size())

	for i := 0; i < dirs.Size(); i++ {
		res.Set(i, FileInfo{Path: dirs.Get(i)})
	}
	return res, err
}

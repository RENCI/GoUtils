package FileSystem

import (
	"os"
	"path"
	"path/filepath"
	"strings"
)

var PATH_SEPARATOR = "/"
var EXTENSION_SEPARATOR = "."

var Path PathObj

type PathObj struct {
}

func Path_New() *PathObj {
	return &PathObj{}
}

func (this *PathObj) Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func (this *PathObj) IsDirectory(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return false, err
	} else {
		if fi.IsDir() {
			return true, nil
		} else {
			return false, nil
		}
	}
}

func (this *PathObj) IsFile(path string) (bool, error) {
	r, err := this.IsDirectory(path)
	return !r, err
}

func (this *PathObj) GetDirectoryParent(path string) string {
	path_to_proc := this.TrimEndingDirectorySeparator(path)
	pos := strings.LastIndex(path_to_proc, PATH_SEPARATOR)
	if pos == -1 {
		return ""
	}
	return path[:pos]
}

func (this *PathObj) TrimEndingDirectorySeparator(path string) string {
	path_to_proc := path

	if path != "" && path[len(path)-1] == PATH_SEPARATOR[0] {
		path_to_proc = path[:len(path)-1]
	}
	return path_to_proc
}

func (this *PathObj) Combine(p1 ...string) string {
	return path.Join(p1...)
}

func (this *PathObj) GetExtension(path string) string {
	path_to_proc := path

	pos_path_sep := strings.LastIndex(path_to_proc, PATH_SEPARATOR)
	pos_ext_sep := strings.LastIndex(path_to_proc, EXTENSION_SEPARATOR)
	if pos_ext_sep == -1 {
		return ""
	}

	if pos_ext_sep < pos_path_sep {
		return ""
	}

	return path[pos_ext_sep:]
}

func (this *PathObj) GetFileName(path string) string {
	path_to_proc := path

	pos_path_sep := strings.LastIndex(path_to_proc, PATH_SEPARATOR)

	if pos_path_sep == -1 {
		return path
	}

	return path[pos_path_sep+1:]
}

func (this *PathObj) GetFullPath(path string) (string, error) {
	absPath, err := filepath.Abs(path)
	return absPath, err
}

func (this *PathObj) GetRelativePath(relativeTo string, path string) (string, error) {
	relativePath, err := filepath.Rel(relativeTo, path)
	return relativePath, err
}

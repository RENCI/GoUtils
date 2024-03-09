package Path

import (
	"os"
	"path"
	"path/filepath"
	"strings"
)

var PATH_SEPARATOR = "/"
var EXTENSION_SEPARATOR = "."

func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func IsDirectory(path string) (bool, error) {
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

func IsFile(path string) (bool, error) {
	r, err := IsDirectory(path)
	return !r, err
}

func GetDirectoryParent(path string) string {
	path_to_proc := TrimEndingDirectorySeparator(path)
	pos := strings.LastIndex(path_to_proc, PATH_SEPARATOR)
	if pos == -1 {
		return ""
	}
	return path[:pos]
}

func TrimEndingDirectorySeparator(path string) string {
	path_to_proc := path

	if path != "" && path[len(path)-1] == PATH_SEPARATOR[0] {
		path_to_proc = path[:len(path)-1]
	}
	return path_to_proc
}

func Combine(p1 ...string) string {
	return path.Join(p1...)
}

func GetExtension(path string) string {
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

func GetFileName(path string) string {
	path_to_proc := path

	pos_path_sep := strings.LastIndex(path_to_proc, PATH_SEPARATOR)

	if pos_path_sep == -1 {
		return path
	}

	return path[pos_path_sep+1:]
}

func GetFullPath(path string) (string, error) {
	absPath, err := filepath.Abs(path)
	return absPath, err
}

func GetRelativePath(relativeTo string, path string) (string, error) {
	relativePath, err := filepath.Rel(relativeTo, path)
	return relativePath, err
}

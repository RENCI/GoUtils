package FileSystem

import (
	"github.com/RENCI/GoUtils/Collections/List"
	"github.com/RENCI/GoUtils/FileSystem/Path"
	"os"
)

func Create(path string) error {
	return os.Mkdir(path, os.ModePerm)
}

func CreateSubdirectiry(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func Delete(path string) error {
	return os.Remove(path)
}

func DeleteAll(path string) error {
	return os.RemoveAll(path)
}

func GetDirectories(path string) (List.List[string], error) {
	files, err := os.ReadDir(path)

	res := List.New[string]()
	for _, v := range files {
		if v.IsDir() {
			res.Add(Path.Combine(path, v.Name()))
		} else {
		}
	}

	return res, err
}

func GetFiles(path string) (List.List[string], error) {
	files, err := os.ReadDir(path)

	res := List.New[string]()
	for _, v := range files {
		if v.IsDir() {

		} else {
			res.Add(Path.Combine(path, v.Name()))
		}
	}

	return res, err
}

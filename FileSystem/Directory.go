package FileSystem

import (
	"github.com/RENCI/GoUtils/Collections"
	"os"
)

var Directory DirectoryObj

type DirectoryObj struct{}

func (this *DirectoryObj) Create(path string) error {
	return os.Mkdir(path, os.ModePerm)
}

func (this *DirectoryObj) CreateSubdirectiry(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func (this *DirectoryObj) Delete(path string) error {
	return os.Remove(path)
}

func (this *DirectoryObj) DeleteAll(path string) error {
	return os.RemoveAll(path)
}

func (this *DirectoryObj) GetDirectories(path string) (Collections.List[string], error) {
	files, err := os.ReadDir(path)

	res := Collections.NewList[string]()
	for _, v := range files {
		if v.IsDir() {
			res.Add(Path.Combine(path, v.Name()))
		} else {
		}
	}

	return res, err
}

func (this *DirectoryObj) GetFiles(path string) (Collections.List[string], error) {
	files, err := os.ReadDir(path)

	res := Collections.NewList[string]()
	for _, v := range files {
		if v.IsDir() {

		} else {
			res.Add(Path.Combine(path, v.Name()))
		}
	}

	return res, err
}

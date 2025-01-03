package FileSystem

import (
	"bufio"
	"encoding/hex"
	"encoding/json"
	"github.com/RENCI/GoUtils/Collections"
	"hash"
	"io"
	"os"
)

type FileInfo struct {
	Path string
}

func FileInfo_New(path string) *FileInfo {
	return &FileInfo{Path: path}
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

func (this *FileInfo) ReadBytes() ([]byte, error) {
	bytes, err := os.ReadFile(this.Path)
	return bytes, err
}

func (this *FileInfo) WriteBytes(data []byte) error {
	err := os.WriteFile(this.Path, data, os.ModePerm)
	return err
}

func (this *FileInfo) ReadAllText() (string, error) {
	bytes, err := os.ReadFile(this.Path)
	return string(bytes), err
}

func (this *FileInfo) WriteAllText(text string) error {
	err := os.WriteFile(this.Path, []byte(text), os.ModePerm)
	return err
}

func (this *FileInfo) ReadFileLines() (Collections.List[string], error) {
	res := Collections.NewList[string]()
	f, err := os.OpenFile(this.Path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return res, err
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				if len(line) > 0 {
					res.Add(line)
				}
				break
			}
			return res, err
		}
		res.Add(line)
	}

	return res, nil
}

func (this *FileInfo) GetFileChecksum(hash hash.Hash) (string, error) {
	//println(filePath)
	//defer timeTrack(time.Now(), fmt.Sprintf("getFileChecksum [%s]: ", filePath))
	file, err := os.Open(this.Path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	hashstr := hex.EncodeToString(hash.Sum(nil))
	return hashstr, nil
}

func (this *FileInfo) LoadJSON() (interface{}, error) {
	var data interface{}
	fileData, err := os.ReadFile(this.Path)
	if err != nil {
		return data, err
	}
	return data, json.Unmarshal(fileData, &data)
}

func (this *FileInfo) Directory() DirectoryInfo {
	res := DirectoryInfo{
		Path: Path.GetDirectoryParent(this.Path),
	}

	return res
}

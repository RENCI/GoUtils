package FileSystem

import (
	"github.com/RENCI/GoUtils/FileSystem"
	"github.com/stretchr/testify/assert"
	"testing"
)

var TEST_PATH_FI = "./Data/File"

func Test_FileInfo_Create(t *testing.T) {

	p := FileSystem.Path.Combine(TEST_PATH_FI, "file1.txt")
	fi := FileSystem.FileInfo_New(p)
	fi.Create()
	exists, err2 := FileSystem.Path.Exists(p)
	assert.True(t, exists)
	assert.Nil(t, err2)
	assert.NotNil(t, fi)

	//// clean up
	err_del := FileSystem.Directory.DeleteAll(p)
	assert.Nil(t, err_del)
}

func Test_FileInfo_Delete(t *testing.T) {

	p := FileSystem.Path.Combine(TEST_PATH_FI, "file2.txt")
	fi := FileSystem.FileInfo_New(p)
	fi.Create()
	exists, err2 := FileSystem.Path.Exists(p)
	assert.True(t, exists)
	assert.Nil(t, err2)
	assert.NotNil(t, fi)

	//// clean up
	err_del := fi.Delete()
	assert.Nil(t, err_del)
}

func Test_FileInfo_Directory(t *testing.T) {

	p := FileSystem.Path.Combine(TEST_PATH_FI, "file2.txt")
	fi := FileSystem.FileInfo_New(p)
	fi.Create()
	exists, err2 := FileSystem.Path.Exists(p)
	assert.True(t, exists)
	assert.Nil(t, err2)

	di := fi.Directory()
	path := di.Path

	assert.Equal(t, len(path), 9)
	//// clean up
	err_del := fi.Delete()
	assert.Nil(t, err_del)
}

func Test_FileInfo_Name(t *testing.T) {

	p := FileSystem.Path.Combine(TEST_PATH_FI, "file2.txt")
	fi := FileSystem.FileInfo_New(p)
	fi.Create()
	name := fi.Name()
	assert.Equal(t, "file2.txt", name)

	di := fi.Directory()
	path := di.Path

	assert.Equal(t, len(path), 9)
	//// clean up
	err_del := fi.Delete()
	assert.Nil(t, err_del)
}

func Test_FileInfo_Exists(t *testing.T) {

	p := FileSystem.Path.Combine(TEST_PATH_FI, "file2.txt")
	fi := FileSystem.FileInfo_New(p)
	fi.Create()
	exists, err2 := fi.Exists()
	assert.True(t, exists)
	assert.Nil(t, err2)
	assert.NotNil(t, fi)

	//// clean up
	err_del := fi.Delete()
	assert.Nil(t, err_del)
}

func Test_FileInfo_ReadAllText(t *testing.T) {

	p := FileSystem.Path.Combine(TEST_PATH_FI, "ReadAllText.txt")
	fi := FileSystem.FileInfo_New(p)
	str, err := fi.ReadAllText()
	assert.Equal(t, "1 2 3 4 5", str)
	assert.Nil(t, err)

}

func Test_FileInfo_ReadBytes(t *testing.T) {

	p := FileSystem.Path.Combine(TEST_PATH_FI, "ReadBytes.txt")
	fi := FileSystem.FileInfo_New(p)
	bs, err := fi.ReadBytes()
	assert.Equal(t, 9, len(bs))
	assert.Nil(t, err)

}

func Test_FileInfo_ReadFileLines(t *testing.T) {

	p := FileSystem.Path.Combine(TEST_PATH_FI, "ReadFileLines.txt")
	fi := FileSystem.FileInfo_New(p)
	str, err := fi.ReadFileLines()
	assert.Equal(t, 3, str.Size())
	assert.Nil(t, err)

}

func Test_FileInfo_WriteAllText(t *testing.T) {

	p := FileSystem.Path.Combine(TEST_PATH_FI, "WriteAllText.txt")
	fi := FileSystem.FileInfo_New(p)
	err := fi.WriteAllText("1 2 3 4 5")

	assert.Nil(t, err)
	str, err2 := fi.ReadAllText()
	assert.Equal(t, "1 2 3 4 5", str)
	assert.Nil(t, err2)

}

func Test_FileInfo_WriteAllBytes(t *testing.T) {

	p := FileSystem.Path.Combine(TEST_PATH_FI, "WriteAllBytes.txt")
	fi := FileSystem.FileInfo_New(p)
	err := fi.WriteBytes([]byte{1, 2, 3, 4, 5})

	assert.Nil(t, err)
	bs, err2 := fi.ReadBytes()
	assert.Equal(t, 5, len(bs))
	assert.Nil(t, err2)

}

func Test_FileInfo_ReadJson(t *testing.T) {

	p := FileSystem.Path.Combine(TEST_PATH_FI, "JsonFile.json")
	fi := FileSystem.FileInfo_New(p)
	obj, err := fi.LoadJSON()

	assert.Nil(t, err)
	assert.NotNil(t, obj)

}

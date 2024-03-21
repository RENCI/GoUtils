package FileSystem

import (
	"github.com/RENCI/GoUtils/FileSystem"
	"github.com/stretchr/testify/assert"
	"testing"
)

var TEST_PATH_FI = "./Data/File"

func Test_FileInfo_Create(t *testing.T) {

	p := fs.Path.Combine(TEST_PATH_FI, "file1.txt")
	fi := FileSystem.FileInfo_New(p)
	fi.Create()
	exists, err2 := fs.Path.Exists(p)
	assert.True(t, exists)
	assert.Nil(t, err2)
	assert.NotNil(t, fi)

	//// clean up
	//err_del := fs.Directory.DeleteAll(p)
	//assert.Nil(t, err_del)
}

package FileSystem

import (
	"github.com/RENCI/GoUtils/FileSystem"
	"github.com/stretchr/testify/assert"
	"testing"
)

var TEST_PATH_DI = "./Data/Directory"

func Test_DirectoryInfo_Create(t *testing.T) {

	p := FileSystem.Path.Combine(TEST_PATH_DI, "NewDir2")
	di := FileSystem.DirectoryInfo_New(p)
	di.Create()
	exists, err2 := FileSystem.Path.Exists(p)
	assert.True(t, exists)
	assert.Nil(t, err2)
	assert.NotNil(t, di)

	// clean up
	err_del := FileSystem.Directory.DeleteAll(p)
	assert.Nil(t, err_del)
}

package FileSystem

import (
	"github.com/RENCI/GoUtils/FileSystem"
	"github.com/stretchr/testify/assert"
	"testing"
)

var TEST_PATH = "./Data/Directory"

func Test_Directory_Create(t *testing.T) {
	p := FileSystem.Path.Combine(TEST_PATH, "NewDir1")
	err := FileSystem.Directory.Create(p)
	exists, err2 := FileSystem.Path.Exists(p)
	assert.True(t, exists)
	assert.Nil(t, err)
	assert.Nil(t, err2)

	// clean up
	err_del := FileSystem.Directory.DeleteAll(p)
	assert.Nil(t, err_del)
}

func Test_Directory_CreateSubdir(t *testing.T) {
	p := FileSystem.Path.Combine(TEST_PATH, "NewDir2/Dir2")
	err := FileSystem.Directory.CreateSubdirectiry(p)
	exists, err2 := FileSystem.Path.Exists(p)
	assert.True(t, exists)
	assert.Nil(t, err)
	assert.Nil(t, err2)

	// clean up
	err_del := FileSystem.Directory.DeleteAll(FileSystem.Path.Combine(TEST_PATH, "NewDir2"))
	assert.Nil(t, err_del)
}

func Test_Directory_GetDirectories_1(t *testing.T) {
	p := FileSystem.Path.Combine(TEST_PATH, "GetDirectories_1")
	res, err := FileSystem.Directory.GetDirectories(p)
	assert.Equal(t, 3, res.Size())
	assert.Nil(t, err)
}

func Test_Directory_GetFiles_1(t *testing.T) {
	p := FileSystem.Path.Combine(TEST_PATH, "GetFiles_1")
	res, err := FileSystem.Directory.GetFiles(p)
	assert.Equal(t, 2, res.Size())
	assert.Nil(t, err)
}

package FileSystem

import (
	"github.com/RENCI/GoUtils/FileSystem"
	"github.com/stretchr/testify/assert"
	"testing"
)

var TEST_PATH = "./Data/Directory"

func Test_Directory_Create(t *testing.T) {
	fs := FileSystem.New()
	p := fs.Path.Combine(TEST_PATH, "NewDir1")
	directory := FileSystem.Directory_new()
	err := directory.Create(p)
	exists, err2 := FileSystem.Path.Exists(p)
	assert.True(t, exists)
	assert.Nil(t, err)
	assert.Nil(t, err2)

	// clean up
	err_del := directory.DeleteAll(p)
	assert.Nil(t, err_del)
}

func Test_Directory_CreateSubdir(t *testing.T) {
	fs := FileSystem.New()
	p := fs.Path.Combine(TEST_PATH, "NewDir2/Dir2")
	directory := FileSystem.Directory_new()
	err := directory.CreateSubdirectiry(p)
	exists, err2 := FileSystem.Path.Exists(p)
	assert.True(t, exists)
	assert.Nil(t, err)
	assert.Nil(t, err2)

	// clean up
	err_del := directory.DeleteAll(fs.Path.Combine(TEST_PATH, "NewDir2"))
	assert.Nil(t, err_del)
}

func Test_Directory_GetDirectories_1(t *testing.T) {
	fs := FileSystem.New()
	p := fs.Path.Combine(TEST_PATH, "GetDirectories_1")
	directory := FileSystem.Directory_new()
	res, err := directory.GetDirectories(p)
	assert.Equal(t, 3, res.Size())
	assert.Nil(t, err)
}

func Test_Directory_GetFiles_1(t *testing.T) {
	fs := FileSystem.New()
	p := fs.Path.Combine(TEST_PATH, "GetFiles_1")
	directory := FileSystem.Directory_new()
	res, err := directory.GetFiles(p)
	assert.Equal(t, 2, res.Size())
	assert.Nil(t, err)
}

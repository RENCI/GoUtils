package FileSystem

import (
	"github.com/RENCI/GoUtils/FileSystem"
	"github.com/stretchr/testify/assert"
	"testing"
)

var TEST_PATH_DI = "Data/Directory"

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

func Test_DirectoryInfo_CreateSubdir(t *testing.T) {

	p := FileSystem.Path.Combine(TEST_PATH_DI, "NewDir2")
	di := FileSystem.DirectoryInfo_New(p)
	di.CreateSubdirectory("NewDirS1")

	s1 := FileSystem.Path.Combine(p, "NewDirS1")
	exists, err2 := FileSystem.Path.Exists(s1)
	assert.True(t, exists)
	assert.Nil(t, err2)
	assert.NotNil(t, di)

	// clean up
	err_del := FileSystem.Directory.DeleteAll(p)
	assert.Nil(t, err_del)
}

func Test_DirectoryInfo_GetAbsPath(t *testing.T) {
	p := "."
	di := FileSystem.DirectoryInfo_New(p)
	absPath, err := di.GetAbsPath()
	assert.True(t, len(absPath) > 1)
	assert.Nil(t, err)
}

func Test_DirectoryInfo_GetDirectories(t *testing.T) {
	p := FileSystem.Path.Combine(TEST_PATH_DI, "GetDirectories_1")
	di := FileSystem.DirectoryInfo_New(p)

	dirs, err2 := di.GetDirectories()

	assert.True(t, dirs.Size() == 3)
	assert.Nil(t, err2)
}

func Test_DirectoryInfo_GetFiles(t *testing.T) {
	p := FileSystem.Path.Combine(TEST_PATH_DI, "GetFiles_1")
	di := FileSystem.DirectoryInfo_New(p)

	fs, err2 := di.GetFiles()

	assert.True(t, fs.Size() == 2)
	assert.Nil(t, err2)
}

func Test_DirectoryInfo_GetParent(t *testing.T) {
	p := FileSystem.Path.Combine(TEST_PATH_DI, "GetFiles_1")
	di := FileSystem.DirectoryInfo_New(p)

	di2 := di.GetParent()
	res := di2.Path
	assert.Equal(t, res, TEST_PATH_DI)
}

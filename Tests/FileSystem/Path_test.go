package FileSystem_test

import (
	"github.com/RENCI/GoUtils/FileSystem"
	"strings"
	"testing"
)
import "github.com/stretchr/testify/assert"

var fs = FileSystem.New()

func Test_Exists(t *testing.T) {
	res, err := fs.Path.Exists("Path_test.go")
	assert.Equal(t, true, res)
	assert.Nil(t, err)
}

func Test_IsDirectory_Neg(t *testing.T) {
	res, err := fs.Path.IsDirectory("Path_test.go")
	assert.Equal(t, false, res)
	assert.Nil(t, err)
}

func Test_IsDirectory_Pos(t *testing.T) {
	res, err := fs.Path.IsDirectory("Data")
	assert.Equal(t, true, res)
	assert.Nil(t, err)
}

func Test_IsFile_Pos(t *testing.T) {
	res, err := fs.Path.IsFile("Path_test.go")
	assert.Equal(t, true, res)
	assert.Nil(t, err)
}

func Test_GetDirectoryName_1(t *testing.T) {
	p1 := "/opt/soft/files/opt.sh"
	dirname := fs.Path.GetDirectoryParent(p1)
	assert.Equal(t, "/opt/soft/files", dirname)
}

func Test_GetDirectoryName_2(t *testing.T) {
	p1 := "/opt/soft/files/"
	dirname := fs.Path.GetDirectoryParent(p1)
	assert.Equal(t, "/opt/soft", dirname)
}

func Test_GetDirectoryName_3(t *testing.T) {
	p1 := "/opt/soft/files"
	dirname := fs.Path.GetDirectoryParent(p1)
	assert.Equal(t, "/opt/soft", dirname)
}

func Test_GetExtension_1(t *testing.T) {
	p1 := "/opt/soft/files/opt.sh"
	ext := fs.Path.GetExtension(p1)
	assert.Equal(t, ".sh", ext)
}

func Test_GetExtension_2(t *testing.T) {
	p1 := "/opt/soft/files/opt.sh/"
	ext := fs.Path.GetExtension(p1)
	assert.Equal(t, "", ext)
}

func Test_GetFileName_1(t *testing.T) {
	p1 := "/opt/soft/files/opt.sh"
	fn := fs.Path.GetFileName(p1)
	assert.Equal(t, "opt.sh", fn)
}

func Test_GetFileName_2(t *testing.T) {
	p1 := "/opt/soft/files/opt.sh/"
	fn := fs.Path.GetExtension(p1)
	assert.Equal(t, "", fn)
}

func Test_GetFullPath(t *testing.T) {
	p1 := "Path_test.go"
	fp, err := fs.Path.GetFullPath(p1)
	assert.Nil(t, err)
	assert.Greater(t, len(fp), len(p1))
	assert.True(t, strings.HasSuffix(fp, p1))
}

func Test_GetRelativePath_1(t *testing.T) {
	rp, err := fs.Path.GetRelativePath("/opt/test", "/opt/Common")
	assert.Nil(t, err)
	assert.Equal(t, "../Common", rp)
}

func Test_GetRelativePath_2(t *testing.T) {
	rp, err := fs.Path.GetRelativePath("/opt/test", "/opt/test/test2")
	assert.Nil(t, err)
	assert.Equal(t, "test2", rp)
}

func Test_TrimEndingDirectorySeparator_1(t *testing.T) {
	p1 := "/opt/soft/files/optsh/"
	p2 := fs.Path.TrimEndingDirectorySeparator(p1)
	assert.Equal(t, "/opt/soft/files/optsh", p2)
}

package FileSystem_test

import (
	"strings"
	"testing"

	"github.com/RENCI/GoUtils/FileSystem"
)
import "github.com/stretchr/testify/assert"

func Test_Exists(t *testing.T) {
	res, err := FileSystem.Exists("Path_test.go")
	assert.Equal(t, true, res)
	assert.Nil(t, err)
}

func Test_IsDirectory_Neg(t *testing.T) {
	res, err := FileSystem.IsDirectory("Path_test.go")
	assert.Equal(t, false, res)
	assert.Nil(t, err)
}

func Test_IsDirectory_Pos(t *testing.T) {
	res, err := FileSystem.IsDirectory("Data")
	assert.Equal(t, true, res)
	assert.Nil(t, err)
}

func Test_IsFile_Pos(t *testing.T) {
	res, err := FileSystem.IsFile("Path_test.go")
	assert.Equal(t, true, res)
	assert.Nil(t, err)
}

func Test_GetDirectoryName_1(t *testing.T) {
	p1 := "/opt/soft/files/opt.sh"
	dirname := FileSystem.GetDirectoryName(p1)
	assert.Equal(t, "/opt/soft/files", dirname)
}

func Test_GetDirectoryName_2(t *testing.T) {
	p1 := "/opt/soft/files/"
	dirname := FileSystem.GetDirectoryName(p1)
	assert.Equal(t, "/opt/soft", dirname)
}

func Test_GetDirectoryName_3(t *testing.T) {
	p1 := "/opt/soft/files"
	dirname := FileSystem.GetDirectoryName(p1)
	assert.Equal(t, "/opt/soft", dirname)
}

func Test_GetExtension_1(t *testing.T) {
	p1 := "/opt/soft/files/opt.sh"
	ext := FileSystem.GetExtension(p1)
	assert.Equal(t, ".sh", ext)
}

func Test_GetExtension_2(t *testing.T) {
	p1 := "/opt/soft/files/opt.sh/"
	ext := FileSystem.GetExtension(p1)
	assert.Equal(t, "", ext)
}

func Test_GetFileName_1(t *testing.T) {
	p1 := "/opt/soft/files/opt.sh"
	fn := FileSystem.GetFileName(p1)
	assert.Equal(t, "opt.sh", fn)
}

func Test_GetFileName_2(t *testing.T) {
	p1 := "/opt/soft/files/opt.sh/"
	fn := FileSystem.GetExtension(p1)
	assert.Equal(t, "", fn)
}

func Test_GetFullPath(t *testing.T) {
	p1 := "Path_test.go"
	fp, err := FileSystem.GetFullPath(p1)
	assert.Nil(t, err)
	assert.Greater(t, len(fp), len(p1))
	assert.True(t, strings.HasSuffix(fp, p1))
}

func Test_GetRelativePath_1(t *testing.T) {
	rp, err := FileSystem.GetRelativePath("/opt/test", "/opt/Common")
	assert.Nil(t, err)
	assert.Equal(t, "../Common", rp)
}

func Test_GetRelativePath_2(t *testing.T) {
	rp, err := FileSystem.GetRelativePath("/opt/test", "/opt/test/test2")
	assert.Nil(t, err)
	assert.Equal(t, "test2", rp)
}

func Test_TrimEndingDirectorySeparator_1(t *testing.T) {
	p1 := "/opt/soft/files/optsh/"
	p2 := FileSystem.TrimEndingDirectorySeparator(p1)
	assert.Equal(t, "/opt/soft/files/optsh", p2)
}

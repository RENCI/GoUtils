package FileSystem

type FileSystem struct {
	Directory DirectoryObj
	Path      PathObj
}

func New() *FileSystem {
	return &FileSystem{}
}

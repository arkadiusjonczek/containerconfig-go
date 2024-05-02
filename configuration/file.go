package configuration

type File struct {
	filepath string
	Configuration
}

func NewFile(filepath string) *File {
	return &File{
		filepath: filepath,
		Configuration: Configuration{
			required: true,
		},
	}
}

func (f *File) GetKey() string {
	return f.filepath
}

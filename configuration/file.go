package configuration

type file struct {
	filepath string
	configuration
}

func NewFile(filepath string) *file {
	return &file{
		filepath: filepath,
		configuration: configuration{
			required: true,
		},
	}
}

func (f *file) GetKey() string {
	return f.filepath
}

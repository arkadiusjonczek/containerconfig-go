package configuration

type env struct {
	key string
	configuration
}

func NewEnv(key string) *env {
	return &env{
		key: key,
		configuration: configuration{
			required: true,
		},
	}
}

func (e *env) GetKey() string {
	return e.key
}

package configuration

type Env struct {
	key string
	Configuration
}

func NewEnv(key string) *Env {
	return &Env{
		key: key,
		Configuration: Configuration{
			required: true,
		},
	}
}

func (e *Env) GetKey() string {
	return e.key
}

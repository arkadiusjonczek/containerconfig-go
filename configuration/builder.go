package configuration

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"syscall"
)

type builder[T any] struct {
	envs  []*env
	files []*file
}

func NewBuilder[T any]() *builder[T] {
	return &builder[T]{}
}

func (c *builder[T]) getEnvs() []*env {
	return c.envs
}

func (c *builder[T]) getFiles() []*file {
	return c.files
}

func (c *builder[T]) addEnv(env *env) *env {
	c.envs = append(c.envs, env)

	return env
}

func (c *builder[T]) addFile(file *file) *file {
	c.files = append(c.files, file)

	return file
}

func (c *builder[T]) Env(key string) *env {
	return c.addEnv(NewEnv(key))
}

func (c *builder[T]) File(filepath string) *file {
	return c.addFile(NewFile(filepath))
}

func (c *builder[T]) Build() (*T, error) {
	configuration := new(T)

	valueOfConfiguration := reflect.ValueOf(*configuration)
	if valueOfConfiguration.Kind() != reflect.Struct {
		return nil, fmt.Errorf("the used configuration type is not a struct")
	}

	valueOfConfigurationPtr := reflect.ValueOf(configuration)

	for _, env := range c.getEnvs() {
		envValue, found := syscall.Getenv(env.GetKey())
		if envValue == "" {
			if env.IsRequired() {
				if !found {
					return nil, fmt.Errorf("environment variable '%s' is required but was not found", env.GetKey())
				} else {
					return nil, fmt.Errorf("environment variable '%s' is required but was empty", env.GetKey())
				}
			} else {
				envValue = env.GetDefaultValue()
			}
		}

		var fieldName string
		if env.GetFieldName() != "" {
			fieldName = env.GetFieldName()
		} else {
			fieldName = env.GetKey()
		}

		field := valueOfConfigurationPtr.Elem().FieldByName(fieldName)

		valid := field.IsValid()
		if !valid {
			return nil, fmt.Errorf("field '%s' not found in struct '%T'", env.GetKey(), *configuration)
		}

		isString := field.Kind() == reflect.String
		if !isString {
			return nil, fmt.Errorf("field '%s' of struct '%T' is not a string", env.GetKey(), *configuration)
		}

		field.SetString(envValue)
	}

	for _, file := range c.getFiles() {
		byteValue, err := os.ReadFile(file.GetKey())
		if err != nil {
			return nil, fmt.Errorf("failed to read file '%s'", file.GetKey())
		}
		value := string(byteValue)
		if value == "" {
			if file.IsRequired() {
				return nil, fmt.Errorf("file '%s' is required but was empty", file.GetKey())
			} else {
				value = file.GetDefaultValue()
			}
		}

		var fieldName string
		if file.GetFieldName() != "" {
			fieldName = file.GetFieldName()
		} else {
			fieldName = filepath.Base(file.GetKey())
		}

		field := valueOfConfigurationPtr.Elem().FieldByName(fieldName)

		valid := field.IsValid()
		if !valid {
			return nil, fmt.Errorf("field '%s' not found in struct '%T'", fieldName, *configuration)
		}

		isString := field.Kind() == reflect.String
		if !isString {
			return nil, fmt.Errorf("field '%s' of struct '%T' is not a string", fieldName, *configuration)
		}

		field.SetString(value)
	}

	return configuration, nil
}

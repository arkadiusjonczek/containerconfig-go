package configuration_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/arkadiusjonczek/containerconfig-go/configuration"
)

type TestConfiguration struct {
	RequiredField string
	OptionalField string
	CustomField   string
	IntegerField  int
	Emptyfile     string
	Testfile      string
}

func Test_ConfigurationFromEnv_Error_WithWrongType(t *testing.T) {
	builder := configuration.NewBuilder[*TestConfiguration]()

	configuration, err := builder.Build()

	require.ErrorContains(t, err, "the used configuration type is not a struct")
	require.Nil(t, configuration)
}

func Test_ConfigurationFromEnv_Error_EnvVarIsRequiredButWasNotFound(t *testing.T) {
	builder := configuration.NewBuilder[TestConfiguration]()
	builder.Env("EnvVarNotExist")

	configuration, err := builder.Build()

	require.ErrorContains(t, err, fmt.Sprintf("environment variable '%s' is required but was not found", "EnvVarNotExist"))
	require.Nil(t, configuration)
}

func Test_ConfigurationFromEnv_Error_EnvVarIsRequiredButWasEmpty(t *testing.T) {
	builder := configuration.NewBuilder[TestConfiguration]()
	builder.Env("EnvVarEmpty")
	os.Setenv("EnvVarEmpty", "")

	configuration, err := builder.Build()

	require.ErrorContains(t, err, fmt.Sprintf("environment variable '%s' is required but was empty", "EnvVarEmpty"))
	require.Nil(t, configuration)
}

func Test_ConfigurationFromEnv_Error_FieldNotFoundInStruct(t *testing.T) {
	builder := configuration.NewBuilder[TestConfiguration]()
	builder.Env("FieldNotExistsInStruct")
	os.Setenv("FieldNotExistsInStruct", "foobar")

	configuration, err := builder.Build()

	require.ErrorContains(t, err, fmt.Sprintf("field '%s' not found in struct '%T'", "FieldNotExistsInStruct", TestConfiguration{}))
	require.Nil(t, configuration)
}

func Test_ConfigurationFromEnv_Error_FieldOfStructIsNotAString(t *testing.T) {
	builder := configuration.NewBuilder[TestConfiguration]()
	builder.Env("IntegerField")
	os.Setenv("IntegerField", "foobar")

	configuration, err := builder.Build()

	require.ErrorContains(t, err, fmt.Sprintf("field '%s' of struct '%T' is not a string", "IntegerField", TestConfiguration{}))
	require.Nil(t, configuration)
}

func Test_WithFile_Error_FailedToReadFile(t *testing.T) {
	builder := configuration.NewBuilder[TestConfiguration]()
	builder.File("testfile")

	configuration, err := builder.Build()

	require.ErrorContains(t, err, fmt.Sprintf("failed to read file '%s'", "testfile"))
	require.Nil(t, configuration)
}

func Test_WithFile_Error_EmptyFile(t *testing.T) {
	builder := configuration.NewBuilder[TestConfiguration]()
	builder.File("../test/emptyfile")

	configuration, err := builder.Build()

	require.ErrorContains(t, err, fmt.Sprintf("file '%s' is required but was empty", "../test/emptyfile"))
	require.Nil(t, configuration)
}

func Test_ConfigurationFromEnv_WithSuccessfulConfiguration(t *testing.T) {
	requiredValue := "foo"
	optionalValue := "bar"
	customFieldValue := "custom"

	os.Setenv("RequiredField", requiredValue)
	os.Setenv("CUSTOM_FIELD", customFieldValue)

	builder := configuration.NewBuilder[TestConfiguration]()
	builder.Env("RequiredField")
	builder.Env("OptionalField").SetOptional().WithDefault(optionalValue)
	builder.Env("CUSTOM_FIELD").UseFieldName("CustomField")

	fileValue := "test"
	optionalFileValue := "egal"

	builder.File("../test/Emptyfile").SetOptional().WithDefault(optionalFileValue)
	builder.File("../test/Testfile")

	configuration, err := builder.Build()

	require.NoError(t, err)
	require.IsType(t, &TestConfiguration{}, configuration)

	require.Equal(t, requiredValue, configuration.RequiredField)
	require.Equal(t, optionalValue, configuration.OptionalField)
	require.Equal(t, customFieldValue, configuration.CustomField)

	require.Equal(t, optionalFileValue, configuration.Emptyfile)
	require.Equal(t, fileValue, configuration.Testfile)
}

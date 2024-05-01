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
	IntegerField  int
}

func Test_ConfigurationFromEnv_Error_WithWrongType(t *testing.T) {
	ruleset := configuration.NewRuleSet()

	configuration, err := configuration.Create[*TestConfiguration](ruleset)

	require.ErrorContains(t, err, "the used configuration type is not a struct")
	require.Nil(t, configuration)
}

func Test_ConfigurationFromEnv_Error_EnvVarIsRequiredButWasNotFound(t *testing.T) {
	ruleset := configuration.NewRuleSet()
	ruleset.Required("EnvVarNotExist")

	configuration, err := configuration.Create[TestConfiguration](ruleset)

	require.ErrorContains(t, err, fmt.Sprintf("environment variable '%s' is required but was not found", "EnvVarNotExist"))
	require.Nil(t, configuration)
}

func Test_ConfigurationFromEnv_Error_EnvVarIsRequiredButWasEmpty(t *testing.T) {
	ruleset := configuration.NewRuleSet()
	ruleset.Required("EnvVarEmpty")
	os.Setenv("EnvVarEmpty", "")

	configuration, err := configuration.Create[TestConfiguration](ruleset)

	require.ErrorContains(t, err, fmt.Sprintf("environment variable '%s' is required but was empty", "EnvVarEmpty"))
	require.Nil(t, configuration)
}

func Test_ConfigurationFromEnv_Error_FieldNotFoundInStruct(t *testing.T) {
	ruleset := configuration.NewRuleSet()
	ruleset.Required("FieldNotExistsInStruct")
	os.Setenv("FieldNotExistsInStruct", "foobar")

	configuration, err := configuration.Create[TestConfiguration](ruleset)

	require.ErrorContains(t, err, fmt.Sprintf("field '%s' not found in struct '%T'", "FieldNotExistsInStruct", TestConfiguration{}))
	require.Nil(t, configuration)
}

func Test_ConfigurationFromEnv_Error_FieldOfStructIsNotAString(t *testing.T) {
	ruleset := configuration.NewRuleSet()
	ruleset.Required("IntegerField")
	os.Setenv("IntegerField", "foobar")

	configuration, err := configuration.Create[TestConfiguration](ruleset)

	require.ErrorContains(t, err, fmt.Sprintf("field '%s' of struct '%T' is not a string", "IntegerField", TestConfiguration{}))
	require.Nil(t, configuration)
}

func Test_ConfigurationFromEnv_WithSuccessfulConfiguration(t *testing.T) {
	requiredValue := "foo"
	optionalValue := "bar"

	os.Setenv("RequiredField", requiredValue)

	ruleset := configuration.NewRuleSet()
	ruleset.Required("RequiredField")
	ruleset.Optional("OptionalField", optionalValue)

	c, err := configuration.Create[TestConfiguration](ruleset)

	require.NoError(t, err)
	require.IsType(t, &TestConfiguration{}, c)
	require.Equal(t, c.RequiredField, requiredValue)
	require.Equal(t, c.OptionalField, optionalValue)
}

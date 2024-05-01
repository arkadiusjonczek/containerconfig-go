package configuration

import (
	"fmt"
	"reflect"
	"syscall"
)

func Create[T any](ruleset *RuleSet) (*T, error) {
	configuration := new(T)

	valueOfConfiguration := reflect.ValueOf(*configuration)
	if valueOfConfiguration.Kind() != reflect.Struct {
		return nil, fmt.Errorf("the used configuration type is not a struct")
	}

	valueOfConfigurationPtr := reflect.ValueOf(configuration)

	for _, rule := range ruleset.GetRules() {
		envValue, found := syscall.Getenv(rule.Key)
		if envValue == "" {
			if rule.Required {
				if !found {
					return nil, fmt.Errorf("environment variable '%s' is required but was not found", rule.Key)
				} else {
					return nil, fmt.Errorf("environment variable '%s' is required but was empty", rule.Key)
				}
			} else {
				envValue = rule.DefaultValue
			}
		}

		field := valueOfConfigurationPtr.Elem().FieldByName(rule.Key)

		valid := field.IsValid()
		if !valid {
			return nil, fmt.Errorf("field '%s' not found in struct '%T'", rule.Key, *configuration)
		}

		isString := field.Kind() == reflect.String
		if !isString {
			return nil, fmt.Errorf("field '%s' of struct '%T' is not a string", rule.Key, *configuration)
		}

		field.SetString(envValue)
	}

	return configuration, nil
}

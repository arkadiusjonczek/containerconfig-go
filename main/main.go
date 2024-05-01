package main

import (
	"log"

	"github.com/arkadiusjonczek/containerconfig-go/configuration"
)

type CustomConfiguration struct {
	RequiredField string
	OptionalField string
}

func main() {
	ruleset := configuration.NewRuleSet()
	ruleset.Required("RequiredField")
	ruleset.Optional("OptionalField", "optional")

	customConfiguration, err := configuration.Create[CustomConfiguration](ruleset)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(customConfiguration.RequiredField)
	log.Println(customConfiguration.OptionalField)
}

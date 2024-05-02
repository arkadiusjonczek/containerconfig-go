package main

import (
	"github.com/arkadiusjonczek/containerconfig-go/configuration"
	"log"
)

type CustomConfiguration struct {
	RequiredField string
	OptionalField string
}

func main() {
	builder := configuration.NewBuilder[CustomConfiguration]()
	builder.Env("RequiredField")
	builder.Env("OptionalField").SetOptional().WithDefault("optional")

	customConfiguration, err := builder.Build()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(customConfiguration.RequiredField)
	log.Println(customConfiguration.OptionalField)
}

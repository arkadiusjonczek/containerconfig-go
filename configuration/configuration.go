package configuration

type Configuration struct {
	required     bool
	defaultValue string
}

func (c *Configuration) SetOptional() *Configuration {
	c.required = false

	return c
}

func (c *Configuration) WithDefault(defaultValue string) *Configuration {
	c.defaultValue = defaultValue

	return c
}

func (c *Configuration) IsRequired() bool {
	return c.required
}

func (c *Configuration) GetDefaultValue() string {
	return c.defaultValue
}

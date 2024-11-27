package configuration

type configuration struct {
	required     bool
	defaultValue string
	fieldName    string
}

func (c *configuration) SetOptional() *configuration {
	c.required = false

	return c
}

func (c *configuration) WithDefault(defaultValue string) *configuration {
	c.defaultValue = defaultValue

	return c
}

func (c *configuration) UseFieldName(fieldName string) *configuration {
	c.fieldName = fieldName

	return c
}

func (c *configuration) IsRequired() bool {
	return c.required
}

func (c *configuration) GetDefaultValue() string {
	return c.defaultValue
}

func (c *configuration) GetFieldName() string {
	return c.fieldName
}

package configuration

type Rule struct {
	Key          string
	DefaultValue string
	Required     bool
}

func NewRule(key string, defaultValue string, required bool) *Rule {
	return &Rule{
		Key:          key,
		DefaultValue: defaultValue,
		Required:     required,
	}
}

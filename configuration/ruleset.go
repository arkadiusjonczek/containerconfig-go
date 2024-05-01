package configuration

type RuleSet struct {
	rules []*Rule
}

func NewRuleSet() *RuleSet {
	return &RuleSet{}
}

func (d *RuleSet) GetRules() []*Rule {
	return d.rules
}

func (d *RuleSet) Required(key string) *RuleSet {
	value := NewRule(key, "", true)
	d.rules = append(d.rules, value)
	return d
}

func (d *RuleSet) Optional(key string, defaultValue string) *RuleSet {
	value := NewRule(key, defaultValue, false)
	d.rules = append(d.rules, value)
	return d
}

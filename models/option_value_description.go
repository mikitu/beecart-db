package models

type OptionValueDescription struct {
	OptionValueId *OptionValue `orm:"column(option_value_id);rel(fk)"`
	LanguageId    *Language    `orm:"column(language_id);rel(fk)"`
	OptionId      *Option      `orm:"column(option_id);rel(fk)"`
	Name          string       `orm:"column(name);size(128)"`
}

package models

type OptionDescription struct {
	OptionId   *Option   `orm:"column(option_id);rel(fk)"`
	LanguageId *Language `orm:"column(language_id);rel(fk)"`
	Name       string    `orm:"column(name);size(128)"`
}

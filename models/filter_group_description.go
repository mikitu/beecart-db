package models

type FilterGroupDescription struct {
	FilterGroupId *FilterGroup `orm:"column(filter_group_id);rel(fk)"`
	LanguageId    *Language    `orm:"column(language_id);rel(fk)"`
	Name          string       `orm:"column(name);size(64)"`
}

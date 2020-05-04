package models

type FilterDescription struct {
	FilterId      *Filter      `orm:"column(filter_id);rel(fk)"`
	LanguageId    *Language    `orm:"column(language_id);rel(fk)"`
	FilterGroupId *FilterGroup `orm:"column(filter_group_id);rel(fk)"`
	Name          string       `orm:"column(name);size(64)"`
}

package models

type CustomFieldDescription struct {
	CustomFieldId *CustomField `orm:"column(custom_field_id);rel(fk)"`
	LanguageId    *Language    `orm:"column(language_id);rel(fk)"`
	Name          string       `orm:"column(name);size(128)"`
}

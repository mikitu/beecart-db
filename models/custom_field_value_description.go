package models

type CustomFieldValueDescription struct {
	CustomFieldValueId *CustomFieldValue `orm:"column(custom_field_value_id);rel(fk)"`
	LanguageId         *Language         `orm:"column(language_id);rel(fk)"`
	CustomFieldId      *CustomField      `orm:"column(custom_field_id);rel(fk)"`
	Name               string            `orm:"column(name);size(128)"`
}

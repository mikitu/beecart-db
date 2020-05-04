package models

type AttributeDescription struct {
	AttributeId *Attribute `orm:"column(attribute_id);rel(fk)"`
	LanguageId  *Language  `orm:"column(language_id);rel(fk)"`
	Name        string     `orm:"column(name);size(64)"`
}

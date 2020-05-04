package models

type AttributeGroupDescription struct {
	AttributeGroupId *AttributeGroup `orm:"column(attribute_group_id);rel(fk)"`
	LanguageId       *Language       `orm:"column(language_id);rel(fk)"`
	Name             string          `orm:"column(name);size(64)"`
}

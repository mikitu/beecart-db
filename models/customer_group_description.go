package models

type CustomerGroupDescription struct {
	CustomerGroupId *CustomerGroup `orm:"column(customer_group_id);rel(fk)"`
	LanguageId      *Language      `orm:"column(language_id);rel(fk)"`
	Name            string         `orm:"column(name);size(32)"`
	Description     string         `orm:"column(description)"`
}

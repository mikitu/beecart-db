package models

type CustomFieldCustomerGroup struct {
	CustomFieldId   int            `orm:"column(custom_field_id)"`
	CustomerGroupId *CustomerGroup `orm:"column(customer_group_id);rel(fk)"`
	Required        int8           `orm:"column(required)"`
}

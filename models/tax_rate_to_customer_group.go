package models

type TaxRateToCustomerGroup struct {
	TaxRateId       *TaxRate       `orm:"column(tax_rate_id);rel(fk)"`
	CustomerGroupId *CustomerGroup `orm:"column(customer_group_id);rel(fk)"`
}

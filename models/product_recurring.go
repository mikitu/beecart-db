package models

type ProductRecurring struct {
	ProductId       *Product       `orm:"column(product_id);rel(fk)"`
	RecurringId     *Recurring     `orm:"column(recurring_id);rel(fk)"`
	CustomerGroupId *CustomerGroup `orm:"column(customer_group_id);rel(fk)"`
}

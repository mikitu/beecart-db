package models

type ProductFilter struct {
	ProductId *Product `orm:"column(product_id);rel(fk)"`
	FilterId  *Filter  `orm:"column(filter_id);rel(fk)"`
}

package models

type ProductRelated struct {
	ProductId *Product `orm:"column(product_id);rel(fk)"`
	RelatedId *Product `orm:"column(related_id);rel(fk)"`
}

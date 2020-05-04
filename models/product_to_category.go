package models

type ProductToCategory struct {
	ProductId  *Product  `orm:"column(product_id);rel(fk)"`
	CategoryId *Category `orm:"column(category_id);rel(fk)"`
}

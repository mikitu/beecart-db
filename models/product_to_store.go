package models

type ProductToStore struct {
	ProductId *Product `orm:"column(product_id);rel(fk)"`
	StoreId   *Store   `orm:"column(store_id);rel(fk)"`
}

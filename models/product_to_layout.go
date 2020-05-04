package models

type ProductToLayout struct {
	ProductId *Product `orm:"column(product_id);rel(fk)"`
	StoreId   *Store   `orm:"column(store_id);rel(fk)"`
	LayoutId  *Layout  `orm:"column(layout_id);rel(fk)"`
}

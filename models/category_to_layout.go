package models

type CategoryToLayout struct {
	CategoryId *Category `orm:"column(category_id);rel(fk)"`
	StoreId    *Store    `orm:"column(store_id);rel(fk)"`
	LayoutId   *Layout   `orm:"column(layout_id);rel(fk)"`
}

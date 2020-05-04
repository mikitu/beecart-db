package models

type CategoryToStore struct {
	CategoryId *Category `orm:"column(category_id);rel(fk)"`
	StoreId    *Store    `orm:"column(store_id);rel(fk)"`
}

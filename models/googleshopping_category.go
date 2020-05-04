package models

type GoogleshoppingCategory struct {
	GoogleProductCategory string    `orm:"column(google_product_category);size(10)"`
	StoreId               *Store    `orm:"column(store_id);rel(fk)"`
	CategoryId            *Category `orm:"column(category_id);rel(fk)"`
}

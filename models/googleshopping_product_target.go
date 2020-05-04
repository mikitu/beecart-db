package models

type GoogleshoppingProductTarget struct {
	ProductId               *Product              `orm:"column(product_id);rel(fk)"`
	StoreId                 *Store                `orm:"column(store_id);rel(fk)"`
	AdvertiseGoogleTargetId *GoogleshoppingTarget `orm:"column(advertise_google_target_id);rel(fk)"`
}

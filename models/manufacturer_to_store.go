package models

type ManufacturerToStore struct {
	ManufacturerId *Manufacturer `orm:"column(manufacturer_id);rel(fk)"`
	StoreId        *Store        `orm:"column(store_id);rel(fk)"`
}

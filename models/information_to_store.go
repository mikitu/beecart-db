package models

type InformationToStore struct {
	InformationId *Information `orm:"column(information_id);rel(fk)"`
	StoreId       *Store       `orm:"column(store_id);rel(fk)"`
}

package models

type InformationToLayout struct {
	InformationId *Information `orm:"column(information_id);rel(fk)"`
	StoreId       *Store       `orm:"column(store_id);rel(fk)"`
	LayoutId      *Layout      `orm:"column(layout_id);rel(fk)"`
}

package models

type OrderStatus struct {
	OrderStatusId int       `orm:"column(order_status_id)"`
	LanguageId    *Language `orm:"column(language_id);rel(fk)"`
	Name          string    `orm:"column(name);size(32)"`
}

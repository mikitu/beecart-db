package models

type StockStatus struct {
	StockStatusId int       `orm:"column(stock_status_id)"`
	LanguageId    *Language `orm:"column(language_id);rel(fk)"`
	Name          string    `orm:"column(name);size(32)"`
}

package models

type CategoryFilter struct {
	CategoryId *Category `orm:"column(category_id);rel(fk)"`
	FilterId   *Filter   `orm:"column(filter_id);rel(fk)"`
}

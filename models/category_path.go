package models

type CategoryPath struct {
	CategoryId *Category `orm:"column(category_id);rel(fk)"`
	PathId     int       `orm:"column(path_id)"`
	Level      int       `orm:"column(level)"`
}

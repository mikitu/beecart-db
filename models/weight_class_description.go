package models

type WeightClassDescription struct {
	WeightClassId *WeightClass `orm:"column(weight_class_id);rel(fk)"`
	LanguageId    *Language    `orm:"column(language_id);rel(fk)"`
	Title         string       `orm:"column(title);size(32)"`
	Unit          string       `orm:"column(unit);size(4)"`
}

package models

type CategoryDescription struct {
	CategoryId      *Category `orm:"column(category_id);rel(fk)"`
	LanguageId      *Language `orm:"column(language_id);rel(fk)"`
	Name            string    `orm:"column(name);size(255)"`
	Description     string    `orm:"column(description)"`
	MetaTitle       string    `orm:"column(meta_title);size(255)"`
	MetaDescription string    `orm:"column(meta_description);size(255)"`
	MetaKeyword     string    `orm:"column(meta_keyword);size(255)"`
}

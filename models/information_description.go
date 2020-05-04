package models

type InformationDescription struct {
	InformationId   *Information `orm:"column(information_id);rel(fk)"`
	LanguageId      *Language    `orm:"column(language_id);rel(fk)"`
	Title           string       `orm:"column(title);size(64)"`
	Description     string       `orm:"column(description)"`
	MetaTitle       string       `orm:"column(meta_title);size(255)"`
	MetaDescription string       `orm:"column(meta_description);size(255)"`
	MetaKeyword     string       `orm:"column(meta_keyword);size(255)"`
}

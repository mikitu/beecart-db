package models

type ProductAttribute struct {
	ProductId   *Product   `orm:"column(product_id);rel(fk)"`
	AttributeId *Attribute `orm:"column(attribute_id);rel(fk)"`
	LanguageId  *Language  `orm:"column(language_id);rel(fk)"`
	Text        string     `orm:"column(text)"`
}

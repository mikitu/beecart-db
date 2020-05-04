package models

type RecurringDescription struct {
	RecurringId *Recurring `orm:"column(recurring_id);rel(fk)"`
	LanguageId  *Language  `orm:"column(language_id);rel(fk)"`
	Name        string     `orm:"column(name);size(255)"`
}

package models

type ReturnStatus struct {
	ReturnStatusId int       `orm:"column(return_status_id)"`
	LanguageId     *Language `orm:"column(language_id);rel(fk)"`
	Name           string    `orm:"column(name);size(32)"`
}

package models

type ReturnReason struct {
	ReturnReasonId int       `orm:"column(return_reason_id)"`
	LanguageId     *Language `orm:"column(language_id);rel(fk)"`
	Name           string    `orm:"column(name);size(128)"`
}

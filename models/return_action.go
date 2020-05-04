package models

type ReturnAction struct {
	ReturnActionId int       `orm:"column(return_action_id)"`
	LanguageId     *Language `orm:"column(language_id);rel(fk)"`
	Name           string    `orm:"column(name);size(64)"`
}

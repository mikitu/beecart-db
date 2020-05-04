package models

type DownloadDescription struct {
	DownloadId *Download `orm:"column(download_id);rel(fk)"`
	LanguageId *Language `orm:"column(language_id);rel(fk)"`
	Name       string    `orm:"column(name);size(64)"`
}

package models

type ProductToDownload struct {
	ProductId  *Product  `orm:"column(product_id);rel(fk)"`
	DownloadId *Download `orm:"column(download_id);rel(fk)"`
}

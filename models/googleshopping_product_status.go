package models

type GoogleshoppingProductStatus struct {
	ProductId            *Product `orm:"column(product_id);rel(fk)"`
	StoreId              *Store   `orm:"column(store_id);rel(fk)"`
	ProductVariationId   string   `orm:"column(product_variation_id);size(64)"`
	DestinationStatuses  string   `orm:"column(destination_statuses)"`
	DataQualityIssues    string   `orm:"column(data_quality_issues)"`
	ItemLevelIssues      string   `orm:"column(item_level_issues)"`
	GoogleExpirationDate int      `orm:"column(google_expiration_date)"`
}

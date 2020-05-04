package models

import "time"

type CustomerWishlist struct {
	CustomerId *Customer `orm:"column(customer_id);rel(fk)"`
	ProductId  *Product  `orm:"column(product_id);rel(fk)"`
	DateAdded  time.Time `orm:"column(date_added);type(datetime)"`
}

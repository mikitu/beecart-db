package models

type CouponCategory struct {
	CouponId   *Coupon   `orm:"column(coupon_id);rel(fk)"`
	CategoryId *Category `orm:"column(category_id);rel(fk)"`
}

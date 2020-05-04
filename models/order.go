package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Order struct {
	Id                    int            `orm:"column(order_id);auto"`
	InvoiceNo             int            `orm:"column(invoice_no)"`
	InvoicePrefix         string         `orm:"column(invoice_prefix);size(26)"`
	StoreId               *Store         `orm:"column(store_id);rel(fk)"`
	StoreName             string         `orm:"column(store_name);size(64)"`
	StoreUrl              string         `orm:"column(store_url);size(255)"`
	CustomerId            *Customer      `orm:"column(customer_id);rel(fk)"`
	CustomerGroupId       *CustomerGroup `orm:"column(customer_group_id);rel(fk)"`
	Firstname             string         `orm:"column(firstname);size(32)"`
	Lastname              string         `orm:"column(lastname);size(32)"`
	Email                 string         `orm:"column(email);size(96)"`
	Telephone             string         `orm:"column(telephone);size(32)"`
	Fax                   string         `orm:"column(fax);size(32)"`
	CustomField           string         `orm:"column(custom_field)"`
	PaymentFirstname      string         `orm:"column(payment_firstname);size(32)"`
	PaymentLastname       string         `orm:"column(payment_lastname);size(32)"`
	PaymentCompany        string         `orm:"column(payment_company);size(40)"`
	PaymentAddress1       string         `orm:"column(payment_address_1);size(128)"`
	PaymentAddress2       string         `orm:"column(payment_address_2);size(128)"`
	PaymentCity           string         `orm:"column(payment_city);size(128)"`
	PaymentPostcode       string         `orm:"column(payment_postcode);size(10)"`
	PaymentCountry        string         `orm:"column(payment_country);size(128)"`
	PaymentCountryId      *Country       `orm:"column(payment_country_id);rel(fk)"`
	PaymentZone           string         `orm:"column(payment_zone);size(128)"`
	PaymentZoneId         *Zone          `orm:"column(payment_zone_id);rel(fk)"`
	PaymentAddressFormat  string         `orm:"column(payment_address_format)"`
	PaymentCustomField    string         `orm:"column(payment_custom_field)"`
	PaymentMethod         string         `orm:"column(payment_method);size(128)"`
	PaymentCode           string         `orm:"column(payment_code);size(128)"`
	ShippingFirstname     string         `orm:"column(shipping_firstname);size(32)"`
	ShippingLastname      string         `orm:"column(shipping_lastname);size(32)"`
	ShippingCompany       string         `orm:"column(shipping_company);size(40)"`
	ShippingAddress1      string         `orm:"column(shipping_address_1);size(128)"`
	ShippingAddress2      string         `orm:"column(shipping_address_2);size(128)"`
	ShippingCity          string         `orm:"column(shipping_city);size(128)"`
	ShippingPostcode      string         `orm:"column(shipping_postcode);size(10)"`
	ShippingCountry       string         `orm:"column(shipping_country);size(128)"`
	ShippingCountryId     *Country       `orm:"column(shipping_country_id);rel(fk)"`
	ShippingZone          string         `orm:"column(shipping_zone);size(128)"`
	ShippingZoneId        *Zone          `orm:"column(shipping_zone_id);rel(fk)"`
	ShippingAddressFormat string         `orm:"column(shipping_address_format)"`
	ShippingCustomField   string         `orm:"column(shipping_custom_field)"`
	ShippingMethod        string         `orm:"column(shipping_method);size(128)"`
	ShippingCode          string         `orm:"column(shipping_code);size(128)"`
	Comment               string         `orm:"column(comment)"`
	Total                 float64        `orm:"column(total);digits(15);decimals(4)"`
	OrderStatusId         int            `orm:"column(order_status_id)"`
	AffiliateId           int            `orm:"column(affiliate_id)"`
	Commission            float64        `orm:"column(commission);digits(15);decimals(4)"`
	MarketingId           *Marketing     `orm:"column(marketing_id);rel(fk)"`
	Tracking              string         `orm:"column(tracking);size(64)"`
	LanguageId            *Language      `orm:"column(language_id);rel(fk)"`
	CurrencyId            *Currency      `orm:"column(currency_id);rel(fk)"`
	CurrencyCode          string         `orm:"column(currency_code);size(3)"`
	CurrencyValue         float64        `orm:"column(currency_value);digits(15);decimals(8)"`
	Ip                    string         `orm:"column(ip);size(40)"`
	ForwardedIp           string         `orm:"column(forwarded_ip);size(40)"`
	UserAgent             string         `orm:"column(user_agent);size(255)"`
	AcceptLanguage        string         `orm:"column(accept_language);size(255)"`
	DateAdded             time.Time      `orm:"column(date_added);type(datetime)"`
	DateModified          time.Time      `orm:"column(date_modified);type(datetime)"`
}

func (t *Order) TableName() string {
	return "order"
}

func init() {
	orm.RegisterModel(new(Order))
}

// AddOrder insert a new Order into database and returns
// last inserted Id on success.
func AddOrder(m *Order) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetOrderById retrieves Order by Id. Returns error if
// Id doesn't exist
func GetOrderById(id int) (v *Order, err error) {
	o := orm.NewOrm()
	v = &Order{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllOrder retrieves all Order matches certain condition. Returns empty list if
// no records exist
func GetAllOrder(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Order))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Order
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateOrder updates Order by Id and returns error if
// the record to be updated doesn't exist
func UpdateOrderById(m *Order) (err error) {
	o := orm.NewOrm()
	v := Order{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteOrder deletes Order by Id and returns error if
// the record to be deleted doesn't exist
func DeleteOrder(id int) (err error) {
	o := orm.NewOrm()
	v := Order{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Order{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

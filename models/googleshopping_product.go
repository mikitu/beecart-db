package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type GoogleshoppingProduct struct {
	Id                    int      `orm:"column(product_advertise_google_id);auto"`
	ProductId             *Product `orm:"column(product_id);rel(fk)"`
	StoreId               *Store   `orm:"column(store_id);rel(fk)"`
	HasIssues             int8     `orm:"column(has_issues);null"`
	DestinationStatus     string   `orm:"column(destination_status)"`
	Impressions           int      `orm:"column(impressions)"`
	Clicks                int      `orm:"column(clicks)"`
	Conversions           int      `orm:"column(conversions)"`
	Cost                  float64  `orm:"column(cost);digits(15);decimals(4)"`
	ConversionValue       float64  `orm:"column(conversion_value);digits(15);decimals(4)"`
	GoogleProductCategory string   `orm:"column(google_product_category);size(10);null"`
	Condition             string   `orm:"column(condition);null"`
	Adult                 int8     `orm:"column(adult);null"`
	Multipack             int      `orm:"column(multipack);null"`
	IsBundle              int8     `orm:"column(is_bundle);null"`
	AgeGroup              string   `orm:"column(age_group);null"`
	Color                 int      `orm:"column(color);null"`
	Gender                string   `orm:"column(gender);null"`
	SizeType              string   `orm:"column(size_type);null"`
	SizeSystem            string   `orm:"column(size_system);null"`
	Size                  int      `orm:"column(size);null"`
	IsModified            int8     `orm:"column(is_modified)"`
}

func (t *GoogleshoppingProduct) TableName() string {
	return "googleshopping_product"
}

func init() {
	orm.RegisterModel(new(GoogleshoppingProduct))
}

// AddGoogleshoppingProduct insert a new GoogleshoppingProduct into database and returns
// last inserted Id on success.
func AddGoogleshoppingProduct(m *GoogleshoppingProduct) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetGoogleshoppingProductById retrieves GoogleshoppingProduct by Id. Returns error if
// Id doesn't exist
func GetGoogleshoppingProductById(id int) (v *GoogleshoppingProduct, err error) {
	o := orm.NewOrm()
	v = &GoogleshoppingProduct{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllGoogleshoppingProduct retrieves all GoogleshoppingProduct matches certain condition. Returns empty list if
// no records exist
func GetAllGoogleshoppingProduct(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GoogleshoppingProduct))
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

	var l []GoogleshoppingProduct
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

// UpdateGoogleshoppingProduct updates GoogleshoppingProduct by Id and returns error if
// the record to be updated doesn't exist
func UpdateGoogleshoppingProductById(m *GoogleshoppingProduct) (err error) {
	o := orm.NewOrm()
	v := GoogleshoppingProduct{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteGoogleshoppingProduct deletes GoogleshoppingProduct by Id and returns error if
// the record to be deleted doesn't exist
func DeleteGoogleshoppingProduct(id int) (err error) {
	o := orm.NewOrm()
	v := GoogleshoppingProduct{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&GoogleshoppingProduct{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

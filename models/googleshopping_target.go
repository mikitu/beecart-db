package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type GoogleshoppingTarget struct {
	Id           int       `orm:"column(advertise_google_target_id);pk"`
	StoreId      *Store    `orm:"column(store_id);rel(fk)"`
	CampaignName string    `orm:"column(campaign_name);size(255)"`
	Country      string    `orm:"column(country);size(2)"`
	Budget       float64   `orm:"column(budget);digits(15);decimals(4)"`
	Feeds        string    `orm:"column(feeds)"`
	Status       string    `orm:"column(status)"`
	DateAdded    time.Time `orm:"column(date_added);type(date);null"`
	Roas         int       `orm:"column(roas)"`
}

func (t *GoogleshoppingTarget) TableName() string {
	return "googleshopping_target"
}

func init() {
	orm.RegisterModel(new(GoogleshoppingTarget))
}

// AddGoogleshoppingTarget insert a new GoogleshoppingTarget into database and returns
// last inserted Id on success.
func AddGoogleshoppingTarget(m *GoogleshoppingTarget) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetGoogleshoppingTargetById retrieves GoogleshoppingTarget by Id. Returns error if
// Id doesn't exist
func GetGoogleshoppingTargetById(id int) (v *GoogleshoppingTarget, err error) {
	o := orm.NewOrm()
	v = &GoogleshoppingTarget{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllGoogleshoppingTarget retrieves all GoogleshoppingTarget matches certain condition. Returns empty list if
// no records exist
func GetAllGoogleshoppingTarget(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GoogleshoppingTarget))
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

	var l []GoogleshoppingTarget
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

// UpdateGoogleshoppingTarget updates GoogleshoppingTarget by Id and returns error if
// the record to be updated doesn't exist
func UpdateGoogleshoppingTargetById(m *GoogleshoppingTarget) (err error) {
	o := orm.NewOrm()
	v := GoogleshoppingTarget{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteGoogleshoppingTarget deletes GoogleshoppingTarget by Id and returns error if
// the record to be deleted doesn't exist
func DeleteGoogleshoppingTarget(id int) (err error) {
	o := orm.NewOrm()
	v := GoogleshoppingTarget{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&GoogleshoppingTarget{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

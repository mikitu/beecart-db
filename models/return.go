package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Return struct {
	Id             int       `orm:"column(return_id);auto"`
	OrderId        *Order    `orm:"column(order_id);rel(fk)"`
	ProductId      *Product  `orm:"column(product_id);rel(fk)"`
	CustomerId     *Customer `orm:"column(customer_id);rel(fk)"`
	Firstname      string    `orm:"column(firstname);size(32)"`
	Lastname       string    `orm:"column(lastname);size(32)"`
	Email          string    `orm:"column(email);size(96)"`
	Telephone      string    `orm:"column(telephone);size(32)"`
	Product        string    `orm:"column(product);size(255)"`
	Model          string    `orm:"column(model);size(64)"`
	Quantity       int       `orm:"column(quantity)"`
	Opened         int8      `orm:"column(opened)"`
	ReturnReasonId int       `orm:"column(return_reason_id)"`
	ReturnActionId int       `orm:"column(return_action_id)"`
	ReturnStatusId int       `orm:"column(return_status_id)"`
	Comment        string    `orm:"column(comment);null"`
	DateOrdered    time.Time `orm:"column(date_ordered);type(date);null"`
	DateAdded      time.Time `orm:"column(date_added);type(datetime)"`
	DateModified   time.Time `orm:"column(date_modified);type(datetime)"`
}

func (t *Return) TableName() string {
	return "return"
}

func init() {
	orm.RegisterModel(new(Return))
}

// AddReturn insert a new Return into database and returns
// last inserted Id on success.
func AddReturn(m *Return) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetReturnById retrieves Return by Id. Returns error if
// Id doesn't exist
func GetReturnById(id int) (v *Return, err error) {
	o := orm.NewOrm()
	v = &Return{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllReturn retrieves all Return matches certain condition. Returns empty list if
// no records exist
func GetAllReturn(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Return))
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

	var l []Return
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

// UpdateReturn updates Return by Id and returns error if
// the record to be updated doesn't exist
func UpdateReturnById(m *Return) (err error) {
	o := orm.NewOrm()
	v := Return{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteReturn deletes Return by Id and returns error if
// the record to be deleted doesn't exist
func DeleteReturn(id int) (err error) {
	o := orm.NewOrm()
	v := Return{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Return{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

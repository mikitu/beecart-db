package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type OrderOption struct {
	Id                   int                 `orm:"column(order_option_id);auto"`
	OrderId              *Order              `orm:"column(order_id);rel(fk)"`
	OrderProductId       *OrderProduct       `orm:"column(order_product_id);rel(fk)"`
	ProductOptionId      *ProductOption      `orm:"column(product_option_id);rel(fk)"`
	ProductOptionValueId *ProductOptionValue `orm:"column(product_option_value_id);rel(fk)"`
	Name                 string              `orm:"column(name);size(255)"`
	Value                string              `orm:"column(value)"`
	Type                 string              `orm:"column(type);size(32)"`
}

func (t *OrderOption) TableName() string {
	return "order_option"
}

func init() {
	orm.RegisterModel(new(OrderOption))
}

// AddOrderOption insert a new OrderOption into database and returns
// last inserted Id on success.
func AddOrderOption(m *OrderOption) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetOrderOptionById retrieves OrderOption by Id. Returns error if
// Id doesn't exist
func GetOrderOptionById(id int) (v *OrderOption, err error) {
	o := orm.NewOrm()
	v = &OrderOption{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllOrderOption retrieves all OrderOption matches certain condition. Returns empty list if
// no records exist
func GetAllOrderOption(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(OrderOption))
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

	var l []OrderOption
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

// UpdateOrderOption updates OrderOption by Id and returns error if
// the record to be updated doesn't exist
func UpdateOrderOptionById(m *OrderOption) (err error) {
	o := orm.NewOrm()
	v := OrderOption{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteOrderOption deletes OrderOption by Id and returns error if
// the record to be deleted doesn't exist
func DeleteOrderOption(id int) (err error) {
	o := orm.NewOrm()
	v := OrderOption{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&OrderOption{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

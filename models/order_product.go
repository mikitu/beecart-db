package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type OrderProduct struct {
	Id        int      `orm:"column(order_product_id);auto"`
	OrderId   *Order   `orm:"column(order_id);rel(fk)"`
	ProductId *Product `orm:"column(product_id);rel(fk)"`
	Name      string   `orm:"column(name);size(255)"`
	Model     string   `orm:"column(model);size(64)"`
	Quantity  int      `orm:"column(quantity)"`
	Price     float64  `orm:"column(price);digits(15);decimals(4)"`
	Total     float64  `orm:"column(total);digits(15);decimals(4)"`
	Tax       float64  `orm:"column(tax);digits(15);decimals(4)"`
	Reward    int      `orm:"column(reward)"`
}

func (t *OrderProduct) TableName() string {
	return "order_product"
}

func init() {
	orm.RegisterModel(new(OrderProduct))
}

// AddOrderProduct insert a new OrderProduct into database and returns
// last inserted Id on success.
func AddOrderProduct(m *OrderProduct) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetOrderProductById retrieves OrderProduct by Id. Returns error if
// Id doesn't exist
func GetOrderProductById(id int) (v *OrderProduct, err error) {
	o := orm.NewOrm()
	v = &OrderProduct{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllOrderProduct retrieves all OrderProduct matches certain condition. Returns empty list if
// no records exist
func GetAllOrderProduct(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(OrderProduct))
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

	var l []OrderProduct
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

// UpdateOrderProduct updates OrderProduct by Id and returns error if
// the record to be updated doesn't exist
func UpdateOrderProductById(m *OrderProduct) (err error) {
	o := orm.NewOrm()
	v := OrderProduct{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteOrderProduct deletes OrderProduct by Id and returns error if
// the record to be deleted doesn't exist
func DeleteOrderProduct(id int) (err error) {
	o := orm.NewOrm()
	v := OrderProduct{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&OrderProduct{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

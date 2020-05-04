package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type OrderRecurring struct {
	Id                   int        `orm:"column(order_recurring_id);auto"`
	OrderId              *Order     `orm:"column(order_id);rel(fk)"`
	Reference            string     `orm:"column(reference);size(255)"`
	ProductId            *Product   `orm:"column(product_id);rel(fk)"`
	ProductName          string     `orm:"column(product_name);size(255)"`
	ProductQuantity      int        `orm:"column(product_quantity)"`
	RecurringId          *Recurring `orm:"column(recurring_id);rel(fk)"`
	RecurringName        string     `orm:"column(recurring_name);size(255)"`
	RecurringDescription string     `orm:"column(recurring_description);size(255)"`
	RecurringFrequency   string     `orm:"column(recurring_frequency);size(25)"`
	RecurringCycle       int16      `orm:"column(recurring_cycle)"`
	RecurringDuration    int16      `orm:"column(recurring_duration)"`
	RecurringPrice       float64    `orm:"column(recurring_price);digits(10);decimals(4)"`
	Trial                int8       `orm:"column(trial)"`
	TrialFrequency       string     `orm:"column(trial_frequency);size(25)"`
	TrialCycle           int16      `orm:"column(trial_cycle)"`
	TrialDuration        int16      `orm:"column(trial_duration)"`
	TrialPrice           float64    `orm:"column(trial_price);digits(10);decimals(4)"`
	Status               int8       `orm:"column(status)"`
	DateAdded            time.Time  `orm:"column(date_added);type(datetime)"`
}

func (t *OrderRecurring) TableName() string {
	return "order_recurring"
}

func init() {
	orm.RegisterModel(new(OrderRecurring))
}

// AddOrderRecurring insert a new OrderRecurring into database and returns
// last inserted Id on success.
func AddOrderRecurring(m *OrderRecurring) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetOrderRecurringById retrieves OrderRecurring by Id. Returns error if
// Id doesn't exist
func GetOrderRecurringById(id int) (v *OrderRecurring, err error) {
	o := orm.NewOrm()
	v = &OrderRecurring{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllOrderRecurring retrieves all OrderRecurring matches certain condition. Returns empty list if
// no records exist
func GetAllOrderRecurring(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(OrderRecurring))
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

	var l []OrderRecurring
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

// UpdateOrderRecurring updates OrderRecurring by Id and returns error if
// the record to be updated doesn't exist
func UpdateOrderRecurringById(m *OrderRecurring) (err error) {
	o := orm.NewOrm()
	v := OrderRecurring{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteOrderRecurring deletes OrderRecurring by Id and returns error if
// the record to be deleted doesn't exist
func DeleteOrderRecurring(id int) (err error) {
	o := orm.NewOrm()
	v := OrderRecurring{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&OrderRecurring{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

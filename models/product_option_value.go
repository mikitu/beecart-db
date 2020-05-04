package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type ProductOptionValue struct {
	Id              int            `orm:"column(product_option_value_id);auto"`
	ProductOptionId *ProductOption `orm:"column(product_option_id);rel(fk)"`
	ProductId       *Product       `orm:"column(product_id);rel(fk)"`
	OptionId        *Option        `orm:"column(option_id);rel(fk)"`
	OptionValueId   *OptionValue   `orm:"column(option_value_id);rel(fk)"`
	Quantity        int            `orm:"column(quantity)"`
	Subtract        int8           `orm:"column(subtract)"`
	Price           float64        `orm:"column(price);digits(15);decimals(4)"`
	PricePrefix     string         `orm:"column(price_prefix);size(1)"`
	Points          int            `orm:"column(points)"`
	PointsPrefix    string         `orm:"column(points_prefix);size(1)"`
	Weight          float64        `orm:"column(weight);digits(15);decimals(8)"`
	WeightPrefix    string         `orm:"column(weight_prefix);size(1)"`
}

func (t *ProductOptionValue) TableName() string {
	return "product_option_value"
}

func init() {
	orm.RegisterModel(new(ProductOptionValue))
}

// AddProductOptionValue insert a new ProductOptionValue into database and returns
// last inserted Id on success.
func AddProductOptionValue(m *ProductOptionValue) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetProductOptionValueById retrieves ProductOptionValue by Id. Returns error if
// Id doesn't exist
func GetProductOptionValueById(id int) (v *ProductOptionValue, err error) {
	o := orm.NewOrm()
	v = &ProductOptionValue{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllProductOptionValue retrieves all ProductOptionValue matches certain condition. Returns empty list if
// no records exist
func GetAllProductOptionValue(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ProductOptionValue))
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

	var l []ProductOptionValue
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

// UpdateProductOptionValue updates ProductOptionValue by Id and returns error if
// the record to be updated doesn't exist
func UpdateProductOptionValueById(m *ProductOptionValue) (err error) {
	o := orm.NewOrm()
	v := ProductOptionValue{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteProductOptionValue deletes ProductOptionValue by Id and returns error if
// the record to be deleted doesn't exist
func DeleteProductOptionValue(id int) (err error) {
	o := orm.NewOrm()
	v := ProductOptionValue{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ProductOptionValue{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Product struct {
	Id             int           `orm:"column(product_id);auto"`
	Model          string        `orm:"column(model);size(64)"`
	Sku            string        `orm:"column(sku);size(64)"`
	Upc            string        `orm:"column(upc);size(12)"`
	Ean            string        `orm:"column(ean);size(14)"`
	Jan            string        `orm:"column(jan);size(13)"`
	Isbn           string        `orm:"column(isbn);size(17)"`
	Mpn            string        `orm:"column(mpn);size(64)"`
	Location       string        `orm:"column(location);size(128)"`
	Quantity       int           `orm:"column(quantity)"`
	StockStatusId  int           `orm:"column(stock_status_id)"`
	Image          string        `orm:"column(image);size(255);null"`
	ManufacturerId *Manufacturer `orm:"column(manufacturer_id);rel(fk)"`
	Shipping       int8          `orm:"column(shipping)"`
	Price          float64       `orm:"column(price);digits(15);decimals(4)"`
	Points         int           `orm:"column(points)"`
	TaxClassId     *TaxClass     `orm:"column(tax_class_id);rel(fk)"`
	DateAvailable  time.Time     `orm:"column(date_available);type(date);null"`
	Weight         float64       `orm:"column(weight);digits(15);decimals(8)"`
	WeightClassId  *WeightClass  `orm:"column(weight_class_id);rel(fk)"`
	Length         float64       `orm:"column(length);digits(15);decimals(8)"`
	Width          float64       `orm:"column(width);digits(15);decimals(8)"`
	Height         float64       `orm:"column(height);digits(15);decimals(8)"`
	LengthClassId  *LengthClass  `orm:"column(length_class_id);rel(fk)"`
	Subtract       int8          `orm:"column(subtract)"`
	Minimum        int           `orm:"column(minimum)"`
	SortOrder      int           `orm:"column(sort_order)"`
	Status         int8          `orm:"column(status)"`
	Viewed         int           `orm:"column(viewed)"`
	DateAdded      time.Time     `orm:"column(date_added);type(datetime)"`
	DateModified   time.Time     `orm:"column(date_modified);type(datetime)"`
}

func (t *Product) TableName() string {
	return "product"
}

func init() {
	orm.RegisterModel(new(Product))
}

// AddProduct insert a new Product into database and returns
// last inserted Id on success.
func AddProduct(m *Product) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetProductById retrieves Product by Id. Returns error if
// Id doesn't exist
func GetProductById(id int) (v *Product, err error) {
	o := orm.NewOrm()
	v = &Product{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllProduct retrieves all Product matches certain condition. Returns empty list if
// no records exist
func GetAllProduct(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Product))
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

	var l []Product
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

// UpdateProduct updates Product by Id and returns error if
// the record to be updated doesn't exist
func UpdateProductById(m *Product) (err error) {
	o := orm.NewOrm()
	v := Product{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteProduct deletes Product by Id and returns error if
// the record to be deleted doesn't exist
func DeleteProduct(id int) (err error) {
	o := orm.NewOrm()
	v := Product{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Product{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

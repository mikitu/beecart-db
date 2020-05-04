package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type CustomerAffiliate struct {
	Id                int       `orm:"column(customer_id);pk"`
	Company           string    `orm:"column(company);size(40)"`
	Website           string    `orm:"column(website);size(255)"`
	Tracking          string    `orm:"column(tracking);size(64)"`
	Commission        float64   `orm:"column(commission);digits(4);decimals(2)"`
	Tax               string    `orm:"column(tax);size(64)"`
	Payment           string    `orm:"column(payment);size(6)"`
	Cheque            string    `orm:"column(cheque);size(100)"`
	Paypal            string    `orm:"column(paypal);size(64)"`
	BankName          string    `orm:"column(bank_name);size(64)"`
	BankBranchNumber  string    `orm:"column(bank_branch_number);size(64)"`
	BankSwiftCode     string    `orm:"column(bank_swift_code);size(64)"`
	BankAccountName   string    `orm:"column(bank_account_name);size(64)"`
	BankAccountNumber string    `orm:"column(bank_account_number);size(64)"`
	CustomField       string    `orm:"column(custom_field)"`
	Status            int8      `orm:"column(status)"`
	DateAdded         time.Time `orm:"column(date_added);type(datetime)"`
}

func (t *CustomerAffiliate) TableName() string {
	return "customer_affiliate"
}

func init() {
	orm.RegisterModel(new(CustomerAffiliate))
}

// AddCustomerAffiliate insert a new CustomerAffiliate into database and returns
// last inserted Id on success.
func AddCustomerAffiliate(m *CustomerAffiliate) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCustomerAffiliateById retrieves CustomerAffiliate by Id. Returns error if
// Id doesn't exist
func GetCustomerAffiliateById(id int) (v *CustomerAffiliate, err error) {
	o := orm.NewOrm()
	v = &CustomerAffiliate{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCustomerAffiliate retrieves all CustomerAffiliate matches certain condition. Returns empty list if
// no records exist
func GetAllCustomerAffiliate(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(CustomerAffiliate))
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

	var l []CustomerAffiliate
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

// UpdateCustomerAffiliate updates CustomerAffiliate by Id and returns error if
// the record to be updated doesn't exist
func UpdateCustomerAffiliateById(m *CustomerAffiliate) (err error) {
	o := orm.NewOrm()
	v := CustomerAffiliate{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCustomerAffiliate deletes CustomerAffiliate by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCustomerAffiliate(id int) (err error) {
	o := orm.NewOrm()
	v := CustomerAffiliate{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&CustomerAffiliate{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

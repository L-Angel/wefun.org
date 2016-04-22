package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type T_all_ip_idc_info struct {
	Id          int
	DeviceId    string `orm:"size(128)"`
	IdcCity     string `orm:"size(128)"`
	IdcOp       string `orm:"size(128)"`
	Location    string `orm:"size(128)"`
	Operation   string `orm:"size(128)"`
	OuterIp     string `orm:"size(128)"`
	Region      string `orm:"size(128)"`
	Zone        string `orm:"size(128)"`
	Subzone     string `orm:"size(128)"`
	Module      string `orm:"size(128)"`
	ModuleOp    string `orm:"size(128)"`
	LanModule   string `orm:"size(128)"`
	WanModule   string `orm:"size(128)"`
	WanModuleOp string `orm:"size(128)"`
	StdOp       string `orm:"size(128)"`
	TimeStamp   string `orm:"size(128)"`
}

func init() {
	orm.RegisterModel(new(T_all_ip_idc_info))
}

// AddT_all_ip_idc_info insert a new T_all_ip_idc_info into database and returns
// last inserted Id on success.
func AddT_all_ip_idc_info(m *T_all_ip_idc_info) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetT_all_ip_idc_infoById retrieves T_all_ip_idc_info by Id. Returns error if
// Id doesn't exist
func GetT_all_ip_idc_infoById(id int64) (v *T_all_ip_idc_info, err error) {
	o := orm.NewOrm()
	v = &T_all_ip_idc_info{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllT_all_ip_idc_info retrieves all T_all_ip_idc_info matches certain condition. Returns empty list if
// no records exist
func GetAllT_all_ip_idc_info(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(T_all_ip_idc_info))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
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

	var l []T_all_ip_idc_info
	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(limit, offset).All(&l, fields...); err == nil {
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

// UpdateT_all_ip_idc_info updates T_all_ip_idc_info by Id and returns error if
// the record to be updated doesn't exist
func UpdateT_all_ip_idc_infoById(m *T_all_ip_idc_info) (err error) {
	o := orm.NewOrm()
	v := T_all_ip_idc_info{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteT_all_ip_idc_info deletes T_all_ip_idc_info by Id and returns error if
// the record to be deleted doesn't exist
func DeleteT_all_ip_idc_info(id int64) (err error) {
	o := orm.NewOrm()
	v := T_all_ip_idc_info{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&T_all_ip_idc_info{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Id            int    `orm:"column(UserId);auto"`
	Nick          string `orm:"column(Nick);size(250);null"`
	Username      string `orm:"column(Username);size(250);null"`
	Password      string `orm:"column(Password);size(250);null"`
	Email         string `orm:"column(Email);size(250);null"`
	Qq            string `orm:"column(Qq);size(50);null"`
	WeChat        string `orm:"column(WeChat);size(50);null"`
	University    string `orm:"column(University);size(50);null"`
	Address       string `orm:"column(Address);size(250);null"`
	Tel           string `orm:"column(Tel);size(50);null"`
	Degree        string `orm:"column(Degree);size(50);null"`
	UniqId        string `orm:"column(UniqId);size(50);null"`
	CommentNotify int    `orm:"column(Notify);default(0)"`
}

func (t *User) TableName() string {
	return "User"
}

func init() {
	orm.RegisterModel(new(User))
}

// AddUser insert a new User into database and returns
// last inserted Id on success.
func AddUser(m *User) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetUserById retrieves User by Id. Returns error if
// Id doesn't exist
func GetUserById(id int) (v *User, err error) {
	o := orm.NewOrm()
	v = &User{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllUser retrieves all User matches certain condition. Returns empty list if
// no records exist
func GetAllUser(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(User))
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

	var l []User
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

// UpdateUser updates User by Id and returns error if
// the record to be updated doesn't exist
func UpdateUserById(m *User) (err error) {
	o := orm.NewOrm()
	v := User{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteUser deletes User by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUser(id int) (err error) {
	o := orm.NewOrm()
	v := User{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&User{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func GetUserByUsername(name string)(v *User,err error){
	o:=orm.NewOrm()
	v = &User{Username:name}
	if err := o.Read(v,"Username");err == nil{
		return v,nil
	}
	return nil,err
}

func CheckOrAddUser(username string,password string,email string,tel string,address string)(r string,msg string,err error){
     o:=orm.NewOrm()
     /**
     err = o.Read(user)
     if err == orm.ErrBoRows {
     	fmt.Println("No this record!")
     } else if err == orm.ErrMissPk{
     	fmt.Println("No Pk")
     }else{
     	o.Insert(user)
     }
     **/
    user := User{Password:password,Username:username,Email:email,Tel:tel,Address:address}
    fmt.Println(user)
    if ok, id, err := o.ReadOrCreate(&user,"Username"); nil == err {
     	fmt.Println(ok)
     	if ok{
     		fmt.Println("Add User Ok")
     		return "true","",nil
     	}else{
     		fmt.Println(id," is existed!")
            return "false","existed",nil
     	}
     }
     fmt.Println(err)
     return "error","",err
}

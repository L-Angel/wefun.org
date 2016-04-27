package models

import (
	"github.com/astaxie/beego/orm"
)

type BlackMan struct {
	Id       int `orm:"column(Id);auto"`
	UserId   int `orm:"column(UserId)"`
	BlackMan int `orm:"column(BlackMan)"`
}

func init() {

	orm.RegisterModel(new(BlackMan))
}

func (t *BlackMan) TableName() string {
	return "BlackMan"
}

func AddBlackMan(b *BlackMan) (id int64, err error) {

	o := orm.NewOrm()
	err = o.Read(b, "UserId", "BlackMan")
	if err == orm.ErrNoRows {
		//fmt.Println(err, "***")
		id, err = o.Insert(b)
	}
	return
}

func DeleteBlackMan(b *BlackMan) (id int64, err error) {
	o := orm.NewOrm()
	if err := o.Read(&b); err == nil {
		id, err = o.Delete(&b)
	}
	return
}

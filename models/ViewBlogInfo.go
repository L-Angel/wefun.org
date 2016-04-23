package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type ViewBlogInfo struct {
	BlogId           int       `orm:"column(BlogId);auto"`
	UserId           int       `orm:"column(UserId);null"`
	Nick             string    `orm:"column(Nick);size(250);null"`
	BlogPermissionId int       `orm:"column(BlogPermissionId);null"`
	PermissionName   string    `orm:"column(PermissionName);size(250);null"`
	SystemTypeId     int       `orm:"column(SystemTypeId);null"`
	SystemTypeName   string    `orm:"column(SystemTypeName);size(250);null"`
	UserTypeId       int       `orm:"column(UserTypeId);null"`
	UserTypeName     string    `orm:"column(UserTypeName);size(250);null"`
	Title            string    `orm:"column(Title);size(250);null"`
	OutLine          string    `orm:"column(OutLine);null"`
	Path             string    `orm:"column(Path);size(512);null"`
	ImgPath          string    `orm:"column(ImgPath);size(512);null"`
	ReleaseTime      time.Time `orm:"column(ReleaseTime);type(datetime);null"`
	ModifyTime       time.Time `orm:"column(ModifyTime);type(datetime);null"`
}

func init() {
	orm.RegisterModel(new(ViewBlogInfo))
}

func (t *ViewBlogInfo) TableName() string {
	return "ViewBlogInfo"
}

func GetAllViewBlogInfo(offset int, limit int) (ml []ViewBlogInfo, err error) {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(new(ViewBlogInfo))
	var l []ViewBlogInfo
	if _, err := qs.Limit(limit, offset).All(&l); err == nil {

		return l, nil
	}
	return nil, err

}

package models

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/astaxie/beego/orm"
)

type Comment struct {
	Id           int       `orm:"column(CommentId);auto"`
	BlogId       int       `orm:"column(BlogId);null"`
	NotifyUserId int       `orm:"column(NotifyUserId);null"`
	UserId       int       `orm:"column(UserId);null"`
	CommentTime  time.Time `orm:"column(CommentTime)"`
	ModifyTime   time.Time `orm:"column(ModifyTime);null"`
	IsAnonymous  bool      `orm:column(IsAnonymous)`
	//Depth       int       `orm:"-"`
	Content string `orm:"column(Content);null"`
	Link    int    `orm:"column(Link)"`
	IsRead  bool   `orm:column(IsRead)`
}

func (t *Comment) TableName() string {
	return "Comment"
}

func init() {
	orm.RegisterModel(new(Comment))
}

// AddComment insert a new Comment into database and returns
// last inserted Id on success.
func AddComment(m *Comment) (id int64, err error) {
	o := orm.NewOrm()

	//Query review authority
	var b []*BlackMan
	qs := o.QueryTable(new(BlackMan))
	qs.Filter("UserId__exact", m.NotifyUserId)

	if _, err = qs.All(&b); err == nil {
		for v, item := range b {
			log.Println(v, "==", item)
			if item.BlackMan == m.UserId {
				return -1, errors.New("Users can not comment on the Blacklist")
			}
		}
		id, err = o.Insert(m)
		if err == nil {
			//There are new comments to notify users
			u := User{Id: m.NotifyUserId}
			o.Read(&u)
			//log.Println(u)
			u.CommentNotify++
			o.Update(&u)
		}
	}
	return
}

// DeleteComment deletes Comment by Id and returns error if
// the record to be deleted doesn't exist
func DeleteComment(id int) (err error) {
	o := orm.NewOrm()
	v := Comment{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		u := User{Id: v.NotifyUserId}
		o.Read(&u)
		log.Println(v)
		if _, err = o.Delete(&v); err == nil && v.IsRead == false && u.CommentNotify > 0 {
			u.CommentNotify -= 1
			log.Println(u)
			_, err = o.Update(&u)
		}
	}
	return
}

// UpdateComment updates Comment by Id and returns error if
// the record to be updated doesn't exist
func UpdateCommentById(m *Comment) (err error) {
	o := orm.NewOrm()
	v := Comment{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// GetCommentById retrieves Comment by Id. Returns error if
// Id doesn't exist
func GetCommentById(id int) (v *Comment, err error) {
	o := orm.NewOrm()
	v = &Comment{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetCommentByNotifyUserId retrieves Comment by NotifyUserId.Return error if UserId doesn't exist
func GetUnReadCommentByNotifyUserId(id int) (comments []*Comment, err error) {
	o := orm.NewOrm()
	//var comments []*Comment

	_, err = o.QueryTable(new(Comment)).Filter("NotifyUserId", id).Filter("IsRead", false).OrderBy("-CommentTime").All(&comments)

	if err == nil {
		return comments, nil
	}
	return nil, err
}

// GetCommentByBlogID retrieves Comment by BlogId.Return error if BlogId doesn't exist
func GetCommentByBlogId(id int) (comments []*Comment, err error) {
	//var comments []*Comment

	o := orm.NewOrm()
	_, err = o.QueryTable(new(Comment)).Filter("BlogId", id).OrderBy("-CommentTime").All(&comments)

	if err == nil {
		return comments, nil
	}
	return nil, err
}

// The reviewer Id find comments
func GetCommentByUserId(id int) (comments []*Comment, err error) {
	//var comments []*Comment

	o := orm.NewOrm()
	_, err = o.QueryTable(new(Comment)).Filter("UserId__exact", id).OrderBy("-CommentTime").All(&comments)
	if err == nil {
		return comments, nil
	}
	return nil, err
}

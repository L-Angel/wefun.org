package tests

import (
	"log"
	//"fmt"
	"testing"
	//	"time"

	"github.com/DeanChina/wefun.org/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	var ConnString = "root:123456@tcp(127.0.0.1:3306)/test"
	beego.Info(ConnString)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", ConnString)
	//orm.RunSyncdb("default", true, true)
}

//func Test_AddComment(t *testing.T) {
//	//	u := models.User{Id: 11, CommentNotify: 0}
//	//	models.AddUser(&u)
//	//用户1将用户10加入黑名单
//	b := models.BlackMan{UserId: 3, BlackMan: 10}
//	models.AddBlackMan(&b)
//	//模拟用户10评论用户1,无权限
//	//id, err := models.AddComment(&models.Comment{2, 1, 3, 10, time.Now(), time.Now(), false, "你好，kitty", 0, false})
//	//用户11，有权限
//	id, err := models.AddComment(&models.Comment{2, 1, 3, 11, time.Now(), time.Now(), false, "你好，kitty", 0, false})
//	log.Println(id, "===", err)
//}

//func Test_DeleteComment(t *testing.T) {
//	err := models.DeleteComment(6)
//	log.Println(err)
//}

//func Test_UpdateCommentById(t *testing.T) {

//	err := models.UpdateCommentById(&models.Comment{10, 1, 3, 10, time.Now(), time.Now(), false, "你好，kittys", 0, false})
//	log.Println(err)
//}

//func Test_GetCommentById(t *testing.T) {
//	v, err := models.GetCommentById(5)
//	if err == nil {
//		log.Println(v)
//	}
//}

//func Test_GetUnReadCommentByNotifyUserId(t *testing.T) {
//	v, err := models.GetUnReadCommentByNotifyUserId(3)
//	if err == nil {
//		for i := 0; i < len(v); i++ {
//			log.Println(v[i])
//		}
//	}
//}

// GetCommentByBlogID retrieves Comment by BlogId.Return error if BlogId doesn't exist
//func Test_GetCommentByBlogId(t *testing.T) {
//	v, err := models.GetCommentByBlogId(1)
//	if err == nil {
//		for i := 0; i < len(v); i++ {
//			log.Println(v[i])
//		}
//	}
//}

// The reviewer Id find comments
//func Test_GetCommentByUserId(t *testing.T) {
//	v, err := models.GetCommentByUserId(10)
//	if err == nil {
//		for i := 0; i < len(v); i++ {
//			log.Println(v[i])
//		}
//	}
//}

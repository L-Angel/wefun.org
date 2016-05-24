package models

type UserGroupHasUser struct {
	UserGroupUserGroupId int   `orm:"column(UserGroup_UserGroupId)"`
	UserUserId           *User `orm:"column(User_UserId);rel(fk)"`
}

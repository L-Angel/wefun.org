package models

type UserRoleHasUser struct {
	UserRoleUserRoleId int   `orm:"column(UserRole_UserRoleId)"`
	UserUserId         *User `orm:"column(User_UserId);rel(fk)"`
}

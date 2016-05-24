package models

type UserAtuthorityDetail struct {
	UserAtuthorityDetailId               int    `orm:"column(UserAtuthorityDetailId)"`
	Name                                 string `orm:"column(Name);size(45);null"`
	AuthorityType                        string `orm:"column(AuthorityType);size(45);null"`
	AuthorityId                          int    `orm:"column(AuthorityId);null"`
	AuthorityDetailAuthorityDetailId     int    `orm:"column(AuthorityDetail_AuthorityDetailId)"`
	AuthorityDescribeAuthorityDescribeId int    `orm:"column(AuthorityDescribe_AuthorityDescribeId)"`
}

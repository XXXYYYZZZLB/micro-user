package model

type User struct {
	//主键
	ID int64 `gorm:"primary_key;not_null;auto_increment"`
	//用户名
	UserName string `grom:"unique_index;not_null"`
	//添加需要的字段
	FirstName string
	//...
	//密码
	HashPassword string
}
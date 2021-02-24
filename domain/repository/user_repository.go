package repository

import (
	"github.com/XXXYYYZZZLB/micro-user/domain/model"
	"github.com/jinzhu/gorm"
)

type IUserRepository interface {
	//初始化数据表
	InitTable() error
	//根据用户名称查找用户信息
	FindUserByName(string)(*model.User ,error)
	//根据用户ID查找用户信息
	FindUserByID(int64)(*model.User , error)
	//创建用户
	CreateUser(*model.User) (int64,error)
	//根据用户ID删除用户
	DeleteUserByID(int64) error
	//更新用户信息
	UpdateUser(*model.User) error
	//查找所有用
	FindAll()([]model.User,error)
}

type UserRepository struct {
	mysqlDb *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{mysqlDb: db}
}

//初始化表
func (u *UserRepository)InitTable() error  {
	return u.mysqlDb.CreateTable(&model.User{}).Error
}

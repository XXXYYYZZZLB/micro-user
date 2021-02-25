package repository

import (
	"github.com/XXXYYYZZZLB/micro-user/domain/model"
	"github.com/jinzhu/gorm"
)

type IUserRepository interface {//接口
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

type UserRepository struct {//相当于类
	mysqlDb *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {//类继承接口？
	return &UserRepository{mysqlDb: db}		//是理解为 为UserRepository赋予接口方法吗？
}

//初始化表
func (u *UserRepository)InitTable() error  {
	return u.mysqlDb.CreateTable(&model.User{}).Error
}

func (u *UserRepository) FindUserByName(name string) (*model.User, error) {
	user := &model.User{}
	return user,u.mysqlDb.Where("user_name = ?",name).Find(user).Error
}

func (u *UserRepository) FindUserByID(userID int64) (*model.User, error) {
	user := &model.User{}
	return user,u.mysqlDb.First(user,userID).Error
}

func (u *UserRepository) CreateUser(user *model.User) (int64, error) {
	return user.ID,u.mysqlDb.Create(user).Error
}

func (u *UserRepository) DeleteUserByID(userID int64) error {
	return u.mysqlDb.Where("id = ?",userID).Delete(&model.User{}).Error
}

func (u *UserRepository) UpdateUser(user *model.User) error {
	return u.mysqlDb.Model(user).Update(&user).Error
}

func (u *UserRepository) FindAll() (userAll []model.User,err error) {
	return userAll,u.mysqlDb.Find(&userAll).Error
}



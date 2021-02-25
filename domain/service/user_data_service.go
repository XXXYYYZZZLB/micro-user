package service

import (
	"github.com/XXXYYYZZZLB/micro-user/domain/model"
	"github.com/XXXYYYZZZLB/micro-user/domain/repository"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/tools/go/ssa/interp/testdata/src/errors"
)

type IUserDataService interface {
	Adduser( *model. User) (int64, error)
	DeleteUser(int64) error
	UpdateUser(user *model.User ,isChangePwd bool)(err error)
	FindUserByName(string)(*model.User,error)
	CheckPwd(userName string , pwd string)(is0k bool,err error)
}

func NewUserDataService(userRepository repository.UserRepository) IUserDataService {
	return &UserDataService{UserRepository:userRepository}
}

type UserDataService struct {
	UserRepository repository.IUserRepository
}

//加密用户密码
func GeneratePassword(userPassword string)([]byte,error)  {
	return bcrypt.GenerateFromPassword([]byte(userPassword),bcrypt.DefaultCost)
}
//
func ValidataPassword(userPassword string,hashed string)(isOk bool,err error){
	if err = bcrypt.CompareHashAndPassword([]byte(hashed),[]byte(userPassword));err != nil{
		return false, errors.New("密码错误")
	}
	return true,nil
}


func (u UserDataService) Adduser(user *model.User) (int64, error) {
	pwdByte,err := GeneratePassword(user.HashPassword)
	if err !=nil{
		return user.ID, err
	}
	user.HashPassword = string(pwdByte)
	return u.UserRepository.CreateUser(user)
}

func (u UserDataService) DeleteUser(userID int64) error {
	return u.UserRepository.DeleteUserByID(userID)
}

func (u UserDataService) UpdateUser(user *model.User, isChangePwd bool) (err error) {
	//判断是否更新了密码
	if isChangePwd{
		pwdByte,err :=GeneratePassword(user.HashPassword)
		if err != nil{
			return err
		}
		user.HashPassword = string(pwdByte)
	}
	//更新日志
	return u.UserRepository.UpdateUser(user)
}

func (u UserDataService) FindUserByName(userName string) (*model.User, error) {
	return u.UserRepository.FindUserByName(userName)
}

func (u UserDataService) CheckPwd(userName string, pwd string) (is0k bool, err error) {
	user,err:=u.UserRepository.FindUserByName(userName)
	if err!=nil{
		return false,err
	}
	return ValidataPassword(pwd,user.HashPassword)
}




package handler

import (
	"context"
	"github.com/XXXYYYZZZLB/micro-user/domain/model"
	"github.com/XXXYYYZZZLB/micro-user/domain/service"
	user "github.com/XXXYYYZZZLB/micro-user/proto/user"
)

type User struct{
	UserDataService service.IUserDataService
}

//注册
func (u *User)Register(ctx context.Context,UserRegisterRequest *user.UserRegisterRequest,UserRegisterResponse *user.UserRegisterResponse) error{
	userRegister :=&model.User{
		UserName: UserRegisterRequest.UserName,
		FirstName: UserRegisterRequest.FirstName,
		HashPassword: UserRegisterRequest.Pwd,
	}
	_,err := u.UserDataService.Adduser(userRegister)
	if err != nil{
		return err
	}
	UserRegisterResponse.Message = "添加成功"
	return nil
}

//登录
func (u *User)Login(cctx context.Context,UserLoginRequest *user.UserLoginRequest,UserUserLoginResponse *user.UserUserLoginResponse) error{
	isOk,err := u.UserDataService.CheckPwd(UserLoginRequest.UserName,UserLoginRequest.Pwd)
	if err != nil{
		return err
	}
	UserUserLoginResponse.IsSuccess = isOk
	return nil
}

//查询用户信息
func (u *User)GetUserInfo(ctx context.Context,UserInfoRequest *user.UserInfoRequest,UserInfoResponse *user.UserInfoResponse) error{
	userInfo,err:=u.UserDataService.FindUserByName(UserInfoRequest.UserName)
	if err!=nil{
		return err
	}
	UserInfoResponse = UserForResponse(userInfo)
	return nil
}


//类型转换
func UserForResponse(userMode *model.User) *user.UserInfoResponse{
	response := &user.UserInfoResponse{}
	response.UserName = userMode.UserName
	response.FirstName = userMode.FirstName
	response.UserId = userMode.ID
	return response
}

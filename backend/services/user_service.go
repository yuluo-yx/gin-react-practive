package services

import (
	"backend/dao"
	"backend/model"
	"backend/pkg/e"
	"backend/pkg/utils"
	"backend/serializer"
	"context"
)

// UserService 接受参数列表
type UserService struct {
	UserName string `json:"user_name" form:"user_name"`
	Password string `json:"password" form:"password"`
}

// Register 用户注册逻辑
func (service UserService) Register(ctx context.Context) serializer.Response {
	var user model.User
	code := e.Success

	// dao层编写
	userDao := dao.NewUserDao(ctx)
	_, exist, err := userDao.ExistOrNotByUserName(service.UserName)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if exist {
		code = e.ErrorExitUser
		return serializer.Response{
			Status: e.ErrorExitUser,
			Msg:    e.GetMsg(code),
		}
	}

	user = model.User{
		UserName: service.UserName,
	}
	//密码加密
	if err = user.SetPassword(service.Password); err != nil {
		code = e.ErrorFailEncryption
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	// 创建用户
	err = userDao.CreateUser(&user)
	if err != nil {
		code = e.Error
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

// Login 用户登录
func (service UserService) Login(ctx context.Context) serializer.Response {
	//定义一个用户对象
	var user *model.User
	// 初始化响应状态码变量值
	code := e.Success
	// 获取userDao对象
	userDao := dao.NewUserDao(ctx)
	//检查用户是否存在
	user, exist, err := userDao.ExistOrNotByUserName(service.UserName)
	if err != nil || !exist {
		code = e.ErrorExitUserNotFound
		// 返回
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "用户不存在！请先注册",
		}
	}
	//检查密码
	if user.CheckPassword(service.Password) == false {
		code = e.ErrorNotCompare
		return serializer.Response{
			Status: e.ErrorNotCompare,
			Msg:    e.GetMsg(code),
			Data:   "密码错误，请重新登录！",
		}
	}
	// 因为http是无状态的协议，所以用户登录过后，让其携带一个登录凭证 使用jwt
	// token 签发
	token, err := utils.GenerateToken(user.ID, service.UserName, 0)
	if err != nil {
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "密码错误，请重新登录！",
		}
	}

	// 返回
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
	}
}

func (service UserService) Get(ctx context.Context) serializer.Response {
	// 初始化响应状态码变量值
	code := e.Success
	// 获取userDao对象
	userDao := dao.NewUserDao(ctx)

	userList, err := userDao.GetUserList()
	if err != nil {
		code = e.ErrorExitUserNotFound
		// 返回
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "获取用户列表信息失败",
		}
	}

	// 返回
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUsers(userList),
	}

}

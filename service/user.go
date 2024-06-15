package service

import (
	"context"
	"register_log/dao"
	"register_log/model"
	"register_log/pkg/e"
	"register_log/pkg/utils"
	"register_log/serializer"
)

type UserService struct {
	NickName string `json:"nick_name" form:"nick_name"`
	UserName string `json:"user_name" form:"user_name"`
	Password string `json:"password" form:"password"`
	Key      string `json:"key" form:"key"`
}

func (service *UserService) Register(ctx context.Context) serializer.Response {
	code := e.Success
	if service.Key == "" || len(service.Key) != 16 {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "密钥长度不足",
		}

	}
	utils.Encrypt.SetKey(service.Key)
	userDao := dao.NewUserDao(ctx)
	_, exist, err := userDao.ExistOrNotByUerName(service.UserName)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if exist {
		code = e.ErrorExistUser
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	user := model.User{
		NickName: service.NickName,
		UserName: service.UserName,
		Avatar:   "avatar.JPG",
		Money:    utils.Encrypt.AesEncoding("100"),
	}
	err = user.SetPassword(service.Password)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	err = userDao.CreateUser(&user)
	if err != nil {
		code = e.Error

	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (service *UserService) Login(ctx context.Context) serializer.Response {
	code := e.Success
	userDao := dao.NewUserDao(ctx)
	user, exist, err := userDao.ExistOrNotByUerName(service.UserName)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if !exist {
		code = e.ErrorUserNotFound
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "用户不存在,请先注册",
		}
	}
	if !user.CheckPassword(service.Password) {
		code = e.ErrorPasswordWrong
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "密码错误,请重新输入",
		}
	}

	token, err := utils.GenerateToken(user.ID, user.NickName, 0)
	if err != nil {
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "密码错误，请重新输入",
		}

	}

	return serializer.Response{
		Status: code,
		Data: serializer.TokenData{
			Token: token,
			User:  serializer.BuildUser(user),
		},
		Msg: e.GetMsg(code),
	}
}

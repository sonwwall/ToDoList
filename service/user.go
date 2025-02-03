package service

import (
	"ToDoList/model"
	"ToDoList/pkg/utils"
	"ToDoList/serializer"
)

type UserService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=15"`
	Password string `form:"password" json:"password" binding:"required,min=5,max=16"`
}

func (service *UserService) Register() serializer.Response {
	var user model.User
	var count int
	model.DB.Model(&model.User{}).Where("user_name=?", service.UserName).
		First(&user).Count(&count)
	if count == 1 {
		return serializer.Response{
			Status: 400,
			Msg:    "用户名已存在",
		}
	}
	user.UserName = service.UserName
	//密码加密
	if err := user.SetPassword(service.Password); err != nil {
		return serializer.Response{
			Status: 400,

			Msg: err.Error(),
		}
	}

	//创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "数据库创建用户失败",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "注册成功",
	}

}

func (service *UserService) Login() serializer.Response {
	var user model.User
	//查询用户
	if err := model.DB.Where("user_name=?", service.UserName).First(&user).Error; err != nil {
		if err.Error() == "record not found" {
			return serializer.Response{
				Status: 400,
				Msg:    "用户不存在",
			}
		}
		return serializer.Response{
			Status: 500,
			Msg:    "数据库查询错误",
		}

	}
	//验证密码
	if user.CheckPassword(service.Password) == false {
		return serializer.Response{
			Status: 400,
			Msg:    "密码错误",
		}
	}
	//发一个token
	token, err := utils.GenerateToken(user.ID, service.UserName, service.Password)
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "token生成失败",
		}
	}
	return serializer.Response{
		Status: 200,
		Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
		Msg:    "登录成功",
	}

}

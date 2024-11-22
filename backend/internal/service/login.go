// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "gf-vue3-admin/api/auth/v1"
)

type (
	ILogin interface {
		// 注册，写入数据库
		Register(ctx context.Context, req *v1.RegisterReq) error
		Lonin(ctx context.Context, req *v1.LoginReq) (string, error)
		// 查询用户
		IsUser(username string) (bool, error)
		IsPasswdCorrect(username string) (bool, error)
	}
)

var (
	localLogin ILogin
)

func Login() ILogin {
	if localLogin == nil {
		panic("implement not found for interface ILogin, forgot register?")
	}
	return localLogin
}

func RegisterLogin(i ILogin) {
	localLogin = i
}

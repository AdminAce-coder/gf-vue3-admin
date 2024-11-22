package auth

import (
	"context"
	modecode "gf-vue3-admin/internal/consts/code/mode"
	"gf-vue3-admin/internal/dao"
	"gf-vue3-admin/internal/model/do"
	"gf-vue3-admin/internal/service"
	"gf-vue3-admin/utility/docode"
)

type sAuth struct{}

func init() {
	service.RegisterAuth(newAuth())
}

func newAuth() *sAuth {
	return &sAuth{}
}

// 查询用户信息

func (s *sAuth) GetUserInfo(ctx context.Context, username string) (*do.User, error) {
	var user *do.User
	err := dao.User.DB().Model("user").
		Where(dao.User.Columns().Username, username).
		Scan(&user)
	if err != nil {
		return nil, docode.NewError(modecode.MODE_SQLFailed, err.Error())
	}
	return user, nil

}

// 查询Codes码

//func (s *sAuth) GetUserCodes(ctx context.Context, username string) (*do.User, error) {
//
//}

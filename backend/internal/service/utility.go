// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"gf-vue3-admin/internal/model/utiliy"
	tl "gf-vue3-admin/utility/terminal"
)

type (
	IUtility interface {
		NewSshSshConn(ctx context.Context, info *utiliy.SshUserInfo) (*tl.SshConn, error)
	}
)

var (
	localUtility IUtility
)

func Utility() IUtility {
	if localUtility == nil {
		panic("implement not found for interface IUtility, forgot register?")
	}
	return localUtility
}

func RegisterUtility(i IUtility) {
	localUtility = i
}

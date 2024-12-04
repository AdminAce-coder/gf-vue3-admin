// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "gf-vue3-admin/api/utility/v1"
)

type (
	IUtility interface {
		// 新增SSH连接信息
		NewSshConnect(ctx context.Context, req *v1.SshUserReq) error
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

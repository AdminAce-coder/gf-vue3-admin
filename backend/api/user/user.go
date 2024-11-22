// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"

	"gf-vue3-admin/api/user/v1"
)

type IUserV1 interface {
	Userinfo(ctx context.Context, req *v1.UserinfoReq) (res *v1.UserinfoRes, err error)
	UserGetCodes(ctx context.Context, req *v1.UserGetCodesReq) (res *v1.UserGetCodesRes, err error)
}

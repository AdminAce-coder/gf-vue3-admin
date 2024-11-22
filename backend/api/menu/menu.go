// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package menu

import (
	"context"

	"gf-vue3-admin/api/menu/v1"
)

type IMenuV1 interface {
	CreateApi(ctx context.Context, req *v1.CreateApiReq) (res *v1.CreateApiRes, err error)
	CreateApiGroup(ctx context.Context, req *v1.CreateApiGroupReq) (res *v1.CreateApiGroupRes, err error)
	DeleteApi(ctx context.Context, req *v1.DeleteApiReq) (res *v1.DeleteApiRes, err error)
}

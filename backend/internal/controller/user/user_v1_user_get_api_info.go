package user

import (
	"context"
	"gf-vue3-admin/api/user/v1"
	"gf-vue3-admin/utility/getApi"
)

func (c *ControllerV1) UserGetApiInfo(ctx context.Context, req *v1.UserGetApiInfoReq) (res *v1.UserGetApiInfoRes, err error) {
	//return nil, gerror.NewCode(gcode.CodeNotImplemented)
	RouteInfo := getApi.GetapiInfo(ctx)
	return &v1.UserGetApiInfoRes{
		ApiInfo: RouteInfo,
	}, nil

}

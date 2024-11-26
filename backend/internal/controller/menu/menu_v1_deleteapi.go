package menu

import (
	"context"
	v1 "gf-vue3-admin/api/menu/v1"
	"gf-vue3-admin/utility/apictrl"
)

func (c *ControllerV1) Deleteapi(ctx context.Context, req *v1.DeleteapiReq) (res *v1.DeleteapiRes, err error) {
	apdle := apictrl.DeleteApi{
		ApiGroup:   req.ApiGroup,
		ApiVersion: req.ApiVersion,
		ApiName:    req.ApiName,
	}
	err = apdle.DeleteApi(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteapiRes{}, nil
}

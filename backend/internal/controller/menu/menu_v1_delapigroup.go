package menu

import (
	"context"
	v1 "gf-vue3-admin/api/menu/v1"
	"gf-vue3-admin/utility/apictrl"
)

func (c *ControllerV1) Delapigroup(ctx context.Context, req *v1.DelapigroupReq) (res *v1.DelapigroupRes, err error) {
	apdel := apictrl.DeleteApiGroup{
		ApiGroupname: req.GroupName,
		Version:      req.Version,
	}
	err = apdel.DeleteGroup(ctx)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

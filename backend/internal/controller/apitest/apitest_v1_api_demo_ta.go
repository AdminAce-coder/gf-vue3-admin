package apitest

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"gf-vue3-admin/api/apitest/v1"
)

func (c *ControllerV1) ApiDemoTa(ctx context.Context, req *v1.ApiDemoTaReq) (res *v1.ApiDemoTaRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

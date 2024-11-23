package menu

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"gf-vue3-admin/api/menu/v1"
)

func (c *ControllerV1) DeleteApi(ctx context.Context, req *v1.DeleteApiReq) (res *v1.DeleteApiRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

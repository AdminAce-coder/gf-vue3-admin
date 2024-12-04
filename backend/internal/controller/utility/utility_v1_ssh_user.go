package utility

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"gf-vue3-admin/api/utility/v1"
)

func (c *ControllerV1) SshUser(ctx context.Context, req *v1.SshUserReq) (res *v1.SshUserRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

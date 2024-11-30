package utility

import (
	"context"
	v1 "gf-vue3-admin/api/utility/v1"
	"gf-vue3-admin/internal/model/utiliy"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

var SshUser *utiliy.SshUserInfo

func (c *ControllerV1) SshInfo(ctx context.Context, req *v1.SshInfoReq) (res *v1.SshInfoRes, err error) {
	// 构造config
	SshUser = &utiliy.SshUserInfo{
		Addr:     req.Addr,
		User:     req.User,
		Password: req.Password,
		Port:     req.Port,
	}
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

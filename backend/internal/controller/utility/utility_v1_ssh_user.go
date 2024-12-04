package utility

import (
	"context"
	"gf-vue3-admin/internal/service"
	"github.com/gogf/gf/v2/os/glog"

	"gf-vue3-admin/api/utility/v1"
)

func (c *ControllerV1) SshUser(ctx context.Context, req *v1.SshUserReq) (res *v1.SshUserRes, err error) {
	//
	err = service.Utility().NewSshConnect(ctx, req)

	if err != nil {
		return nil, err
	}
	glog.Infof(ctx, "SshUser 添加成功")
	return nil, nil
}

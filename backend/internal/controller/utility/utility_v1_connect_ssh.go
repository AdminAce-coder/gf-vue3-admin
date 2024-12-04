package utility

import (
	"context"

	v1 "gf-vue3-admin/api/utility/v1"
	"gf-vue3-admin/internal/model/utiliy"
	"gf-vue3-admin/internal/service"
)

func (c *ControllerV1) ConnectSsh(ctx context.Context, req *v1.ConnectSshReq) (res *v1.ConnectSshRes, err error) {
	// 查询信息
	sshUserInfo, err := service.Utility().SshConnect(ctx, &utiliy.SshConnectInfoInput{
		Host: req.Host,
	})
	if err != nil {
		return nil, err
	}
	// 把查询的数据写入SshUser
	utiliy.SshUser = sshUserInfo
	return nil, nil
}

package utility

import (
	"context"
	"fmt"
	v1 "gf-vue3-admin/api/utility/v1"
	"gf-vue3-admin/internal/model/utiliy"
	"gf-vue3-admin/utility/docode"

	"github.com/gogf/gf/v2/net/gssh"
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

	// 测试SSH连接
	ssh := gssh.SshConfig{
		Userinfo: SshUser,
	}

	// 尝试建立SSH连接
	sshClient, err := ssh.NewSshConfig(ctx)
	if err != nil {
		return nil, docode.NewError(400, fmt.Sprintf("SSH连接失败:%v", err))
	}
	defer sshClient.Client.Close()

	return nil, nil
}

package utility

import (
	"context"
	"fmt"
	v1 "gf-vue3-admin/api/utility/v1"
	"gf-vue3-admin/internal/model/utiliy"
	"gf-vue3-admin/utility/docode"

	gossh "golang.org/x/crypto/ssh"
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
	sshConfig := &gossh.ClientConfig{
		User: req.User,
		Auth: []gossh.AuthMethod{
			gossh.Password(req.Password),
		},
		HostKeyCallback: gossh.InsecureIgnoreHostKey(),
	}
	_, err = gossh.Dial("tcp", fmt.Sprintf("%s:%d", req.Addr, req.Port), sshConfig)
	if err != nil {
		return nil, docode.NewError(400, fmt.Sprintf("SSH连接失败:%v", err))
	}

	return nil, nil
}

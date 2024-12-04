package utility

import (
	"context"
	"fmt"
	"gf-vue3-admin/internal/model/utiliy"
	"gf-vue3-admin/internal/service"
	"gf-vue3-admin/utility/docode"

	gossh "golang.org/x/crypto/ssh"

	v1 "gf-vue3-admin/api/utility/v1"
)

func (c *ControllerV1) SshTry(ctx context.Context, req *v1.SshTryReq) (res *v1.SshTryRes, err error) {
	sshUserInfo, err := service.Utility().SshConnect(ctx, &utiliy.SshConnectInfoInput{
		Host: req.Host,
	})
	if err != nil {
		return nil, err
	}
	fmt.Println("sshUserInfo是%v", sshUserInfo)
	// 测试连接
	// 测试SSH连接
	sshConfig := &gossh.ClientConfig{
		User: sshUserInfo.User,
		Auth: []gossh.AuthMethod{
			gossh.Password(sshUserInfo.Password),
		},
		HostKeyCallback: gossh.InsecureIgnoreHostKey(),
	}
	_, err = gossh.Dial("tcp", fmt.Sprintf("%s:%d", sshUserInfo.Addr, sshUserInfo.Port), sshConfig)
	if err != nil {
		return nil, docode.NewError(400, fmt.Sprintf("SSH连接失败:%v", err))
	}

	return nil, nil

}

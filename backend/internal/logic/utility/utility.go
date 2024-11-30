package utility

import (
	"context"
	"gf-vue3-admin/internal/model/utiliy"
	"gf-vue3-admin/internal/service"
	tl "gf-vue3-admin/utility/terminal"
)

type sUtility struct{}

func init() {
	service.RegisterUtility(new())
}

func new() *sUtility {
	return &sUtility{}
}

func (s *sUtility) NewSshSshConn(ctx context.Context, info *utiliy.SshUserInfo) (*tl.SshConn, error) {

	////Sshconfig
	//tl := tl.Sshconfig{
	//	Userinfo: info,
	//}
	//sshclinet, err := tl.NewSshConfig(ctx)
	//if err != nil {
	//	return nil, err
	//	glog.Error(ctx, "创建sshclinet错误")
	//}
	////NewSshConn 创建 SSH 连接
	//sshConn, err := sshclinet.NewSshConn(2048, 2048)
	//if err != nil {
	//	return nil, err
	//	glog.Error(ctx, "sshConn创建错误")
	//}
	return nil, nil

}

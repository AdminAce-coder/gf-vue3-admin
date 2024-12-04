package utility

import (
	"context"
	v1 "gf-vue3-admin/api/utility/v1"
	"gf-vue3-admin/internal/dao"
	"gf-vue3-admin/internal/model/do"
	"gf-vue3-admin/internal/service"
)

type sUtility struct{}

func init() {
	service.RegisterUtility(new())
}

func new() *sUtility {
	return &sUtility{}
}

// 新增SSH连接信息
func (s *sUtility) NewSshConnect(ctx context.Context, req *v1.SshUserReq) error {
	db := dao.Ssh.DB()
	_, err := db.Model("ssh").Data(do.Ssh{
		HostName: req.Hostname,
		User:     req.User,
		Password: req.Password,
		Port:     req.Port,
		Host:     req.Host,
	}).Insert()
	if err != nil {
		return err
	}
	return nil
}

//

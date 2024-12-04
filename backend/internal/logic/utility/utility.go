package utility

import (
	"context"
	"fmt"
	v1 "gf-vue3-admin/api/utility/v1"
	"gf-vue3-admin/internal/dao"
	"gf-vue3-admin/internal/model/do"
	"gf-vue3-admin/internal/model/utiliy"
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
		Addr:     req.Host,
	}).Insert()
	if err != nil {
		return err
	}
	return nil
}

// 查询连接信息
func (s *sUtility) SshConnect(ctx context.Context, input *utiliy.SshConnectInfoInput) (info *utiliy.SshUserInfo, err error) {
	db := dao.Ssh.DB()
	err = db.Model("ssh").Where(dao.Ssh.Columns().Addr, input.Host).Scan(&info)
	if err != nil {
		return nil, err
	}
	fmt.Println("查询的数据是：%v", info)
	return info, nil
}

// 查询所有信息
func (s *sUtility) GetAllSshinfo(ctx context.Context) ([]utiliy.SshUserInfo, error) {
	var info []utiliy.SshUserInfo
	db := dao.Ssh.DB()
	err := db.Model("ssh").Scan(&info)
	if err != nil {
		return nil, err
	}
	return info, nil
}

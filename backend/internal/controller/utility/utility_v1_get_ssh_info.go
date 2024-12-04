package utility

import (
	"context"
	v1 "gf-vue3-admin/api/utility/v1"
	"gf-vue3-admin/internal/service"
)

func (c *ControllerV1) GetSshInfo(ctx context.Context, req *v1.GetSshInfoReq) (res *v1.GetSshInfoRes, err error) {
	infoList, err := service.Utility().GetAllSshinfo(ctx)
	if err != nil {
		return nil, err
	}

	res = &v1.GetSshInfoRes{
		SshInfoList: make([]struct {
			HostName string `json:"hostname"`
			User     string `json:"user"`
			Port     int    `json:"port"`
			Password string `json:"password"`
			Host     string `json:"host"`
		}, len(infoList)),
	}

	// 将切片数据复制到响应结构中
	for i, info := range infoList {
		res.SshInfoList[i] = struct {
			HostName string `json:"hostname"`
			User     string `json:"user"`
			Port     int    `json:"port"`
			Password string `json:"password"`
			Host     string `json:"host"`
		}{
			HostName: info.HostName,
			User:     info.User,
			Port:     info.Port,
			Password: info.Password,
			Host:     info.Addr,
		}
	}
	return
}

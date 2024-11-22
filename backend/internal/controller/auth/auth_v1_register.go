package auth

import (
	"context"
	v1 "gf-vue3-admin/api/auth/v1"
	"gf-vue3-admin/internal/service"
)

func (c *ControllerV1) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error) {

	err = service.Login().Register(ctx, req)
	if err != nil {
		return nil, err
	}
	return &v1.RegisterRes{
		Message: "注册成功",
	}, nil
}

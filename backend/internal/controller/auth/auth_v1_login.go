package auth

import (
	"context"
	"gf-vue3-admin/api/auth/v1"
	"gf-vue3-admin/internal/service"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	token, err := service.Login().Lonin(ctx, req)
	if err != nil {
		return nil, err
	}

	return &v1.LoginRes{
		AccessToken: token,
	}, nil
}

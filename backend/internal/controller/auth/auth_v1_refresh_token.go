package auth

import (
	"context"
	"gf-vue3-admin/api/auth/v1"
	"gf-vue3-admin/utility/jwt"
	"net/http"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
)

func (c *ControllerV1) RefreshToken(ctx context.Context, req *v1.RefreshTokenReq) (res *v1.RefreshTokenRes, err error) {
	r := g.RequestFromCtx(ctx)
	Bearertoken := r.GetHeader("Authorization")
	//分离Bearer
	parts := strings.Split(Bearertoken, " ")
	if len(parts) < 2 {
		glog.Errorf(ctx, "Invalid Bearer token format")
		r.Response.Write("Token错误:")
	}
	TokenWithoutBearer := strings.TrimSpace(parts[1])
	if Bearertoken == " " {
		glog.Errorf(ctx, "Bearertoken token is empty")
		r.Response.WriteStatus(http.StatusUnauthorized)
	}
	nwetoken, err := jwt.RefreshToken(TokenWithoutBearer)
	if err != nil {
		return nil, err
	}
	return &v1.RefreshTokenRes{
		NewAccessToken: nwetoken,
	}, nil
}

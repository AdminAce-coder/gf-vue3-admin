package user

import (
	"context"
	v1 "gf-vue3-admin/api/user/v1"
	"gf-vue3-admin/internal/service"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

func (c *ControllerV1) Userinfo(ctx context.Context, req *v1.UserinfoReq) (res *v1.UserinfoRes, err error) {
	username := g.RequestFromCtx(ctx).GetCtxVar("username").String()
	user, err := service.Auth().GetUserInfo(ctx, username)
	if err != nil {
		return nil, err
	}

	roles := strings.Split(gconv.String(user.Roles), ",")
	return &v1.UserinfoRes{
		Roles:    roles,
		RealName: gconv.String(user.Realname),
		Username: username,
	}, nil
}

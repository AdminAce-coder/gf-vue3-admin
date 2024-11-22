package user

import (
	"context"
	"gf-vue3-admin/api/user/v1"
	"gf-vue3-admin/internal/service"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

func (c *ControllerV1) UserGetCodes(ctx context.Context, req *v1.UserGetCodesReq) (res *v1.UserGetCodesRes, err error) {
	username := g.RequestFromCtx(ctx).GetCtxVar("username").String()
	user, err := service.Auth().GetUserInfo(ctx, username)
	if err != nil {
		return nil, err
	}

	codes := strings.Split(gconv.String(user.Codes), ",")
	return &v1.UserGetCodesRes{
		Codes: codes,
	}, nil

}

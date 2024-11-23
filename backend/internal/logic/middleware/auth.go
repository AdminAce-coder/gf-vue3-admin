package middleware

import (
	"gf-vue3-admin/internal/service"
	"gf-vue3-admin/utility/jwt"
	"net/http"
	"strings"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
)

type sAuthMiddleware struct{}

func init() {
	service.RegisterAuthMiddleware(newsAuthMiddleware())
}

func newsAuthMiddleware() *sAuthMiddleware {
	return &sAuthMiddleware{}
}

// @ 鉴权中间件

func (s *sAuthMiddleware) AuthMiddleware(r *ghttp.Request) {
	// 判断是否是登录状态---是否有token

	ctx := r.Context()
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

	// 卸载token
	MyClaims, err := jwt.ParseToken(TokenWithoutBearer)
	if err != nil {
		// 设置错误状态和401未授权状态码
		// r.Response.WriteStatus(http.StatusUnauthorized)
		if err.Error() == "token is expired" {
			r.SetError(err)
		} else {
			r.SetError(err)
		}
		return // 直接返回，让 Response 中间件处理响应
	}

	glog.New().Infof(ctx, "解析后的密码是:%s", MyClaims.Password)
	// 验证密码
	ok, err := service.Login().IsPasswdCorrect(MyClaims.Username)
	if err != nil {
		r.Response.Write(err.Error())
	}
	if ok {
		glog.New().Infof(ctx, "鉴权通过:")
		// 将用户信息存储到请求上下文中
		r.SetCtxVar("username", MyClaims.Username)
		r.SetCtxVar("password", MyClaims.Password)
	}
	r.Middleware.Next()
}

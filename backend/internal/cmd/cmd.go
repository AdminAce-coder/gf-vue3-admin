package cmd

import (
	"context"
	"gf-vue3-admin/internal/controller/auth"
	"gf-vue3-admin/internal/controller/user"
	"gf-vue3-admin/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			// auth组 不需要鉴权
			s.Group("/api/auth", func(group *ghttp.RouterGroup) {
				group.Middleware(
					ghttp.MiddlewareHandlerResponse,
					ghttp.MiddlewareCORS,
					service.Middleware().Returndata, // 统一返回数据中间件
				)
				group.Bind(
					auth.NewV1(),
				)
			})
			// user组，需要鉴权
			s.Group("/api/user", func(group *ghttp.RouterGroup) {
				group.Middleware(
					ghttp.MiddlewareHandlerResponse,
					ghttp.MiddlewareCORS,
					service.Middleware().Returndata, // 统一返回数据中间件
					service.AuthMiddleware().AuthMiddleware,
				)
				group.Bind(
					user.NewV1(),
				)
			})

			// 开启 Swagger
			s.SetPort(5321)
			s.Run()
			return nil
		},
	}
)

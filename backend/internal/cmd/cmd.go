package cmd

import (
	"context"
	"gf-vue3-admin/internal/service"
	"gf-vue3-admin/internal/service/register"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			// 动态生成路由组
			for _, route := range register.RouteGroups {
				for path, config := range route {
					if !config.Enable {
						continue
					}
					s.Group(path, func(group *ghttp.RouterGroup) {
						middlewares := []ghttp.HandlerFunc{
							ghttp.MiddlewareHandlerResponse,
							ghttp.MiddlewareCORS,
							service.Middleware().Returndata,
						}

						if config.NeedAuth {
							middlewares = append(middlewares, service.AuthMiddleware().AuthMiddleware)
						}

						group.Middleware(middlewares...)
						if config.GroupName != "" {
							ctrl := register.Get(config.GroupName)
							glog.Infof(ctx, "获取到控制器: %+s\n", config.GroupName)
							if ctrl != nil {
								// 直接绑定控制器实例
								group.Bind(ctrl)
							} else {
								glog.Infof(ctx, "警告: 未找到控制器 %s\n", config.GroupName)
							}
						}
					})
				}
			}
			s.SetPort(5321)
			s.Run()
			return nil
		},
	}
)

// 加载路由配置
func init() {
	err := register.LoadRouteConfig()
	if err != nil {
		glog.Errorf(gctx.New(), "加载配置文件失败%s", err)
	}
}

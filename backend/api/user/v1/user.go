package v1

import (
	"gf-vue3-admin/utility/getApi"

	"github.com/gogf/gf/v2/frame/g"
)

// 带验证的需要验证token
type UserinfoReq struct {
	g.Meta `path:"/info" method:"get" tags:"user" dc:"登录后的用户信息"`
}
type UserinfoRes struct {
	Roles    []string `json:"roles" dc:"权限集" `
	RealName string   `json:"realName" dc:"角色名" `
	Username string   `json:"username" dc:"用户名"`
}

type UserGetCodesReq struct {
	g.Meta `path:"/codes" method:"get" tags:"user" dc:"获取权限码"`
}
type UserGetCodesRes struct {
	Codes []string `json:"codes" dc:"权限码""`
}

// 获取接口信息

type UserGetApiInfoReq struct {
	g.Meta `path:"/apiinfo" method:"get" tags:"user" dc:"获取API信息"`
}
type UserGetApiInfoRes struct {
	ApiInfo *[]getApi.RouteInfo `json:"apiInfo" dc:"权限码""`
}

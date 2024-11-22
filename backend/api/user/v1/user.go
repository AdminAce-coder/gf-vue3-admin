package v1

import "github.com/gogf/gf/v2/frame/g"

// 带验证的需要验证token
type UserinfoReq struct {
	g.Meta `path:"/info" method:"get" tags:"user"`
}
type UserinfoRes struct {
	Roles    []string `json:"roles" dc:"权限集" `
	RealName string   `json:"realName" dc:"角色名" `
	Username string   `json:"username" dc:"用户名"`
}

type UserGetCodesReq struct {
	g.Meta `path:"/codes" method:"get" tags:"user"`
}
type UserGetCodesRes struct {
	Codes []string `json:"codes" dc:"权限码""`
}

package v1

import "github.com/gogf/gf/v2/frame/g"

// 登录
type LoginReq struct {
	g.Meta        `path:"/login" method:"post" tags:"login" dc:"登录接口"`
	SelectAccount string `v:"required" json:"selectAccount" dc:"用户组"  `
	Username      string `v:"required" json:"username" dc:"用户名"  `
	Password      string `v:"required" json:"password"  dc:"密码 " `
}
type LoginRes struct {
	AccessToken string `json:"accessToken" dc:"返回token"`
}

// 注册 不需要鉴权
type RegisterReq struct {
	g.Meta    `path:"/register" method:"post" tags:"login" dc:"注册接口"`
	Username  string `v:"required" json:"username" dc:"用户名"`
	Password  string `v:"required" json:"password"  dc:"密码" `
	Password2 string `v:"required|ci|same:Password" json:"password2"`
}
type RegisterRes struct {
	Message string `json:"message" dc:"消息"`
}

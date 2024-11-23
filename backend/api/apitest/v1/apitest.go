package v1

import "github.com/gogf/gf/v2/frame/g"

type CreateApiReq struct {
	g.Meta `path:"createapi" method:"post" tags:"menu" dc:"菜单相关"`
	// 名称
	name string `v:"required" dc:"名称"`
	// 年龄
	age int `v:"required" dc:"年龄"`
}

type CreateApiRes struct {
}

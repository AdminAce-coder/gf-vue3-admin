package v1

import "github.com/gogf/gf/v2/frame/g"

type ApiDsdsdseleteReq struct {
	g.Meta `path:"/apidsdsdselete" method:"post" tags:"apitest" dc:"菜单相关"`
	// 名称
	name string `v:"required" dc:"名称"`
	// 年龄
	age int `v:"required" dc:"年龄"`
}

type ApiDsdsdseleteRes struct {
}

type ApiDsdsadsvvReq struct {
	g.Meta `path:"/apidsdsadsvv" method:"post" tags:"apitest" dc:"菜单相关"`
	// 名称
	name string `v:"required" dc:"名称"`
	// 年龄
	age int `v:"required" dc:"年龄"`
}

type ApiDsdsadsvvRes struct {
}

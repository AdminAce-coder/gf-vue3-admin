package v1

import "github.com/gogf/gf/v2/frame/g"

// 系统管理相关API
type Parameter struct {
	ParameterName string `json:"parametername"` // 参数名
	DataType      string `json:"datatype"`      // 数据类型
	Required      bool   `json:"required"`
	Description   string `json:"description"` // 描述
}

type CreateApiReq struct {
	g.Meta           `path:"/createapi" method:"post" tags:"menu" dc:"新增API"`
	ApiName          string      `json:"apiName"`
	Isauthentication bool        `json:"isauthentication"` // 是否鉴权
	ApiVersion       string      `json:"apiversion"`
	Method           string      `json:"method"`
	ApiGroup         string      `json:"apiGroup"`
	Description      string      `json:"description"`
	Parameters       []Parameter `json:"parameters"` // 参数
}
type CreateApiRes struct {
}

type CreateApiGroupReq struct {
	g.Meta       `path:"/capigp" method:"post" tags:"menu" dc:"新增分组"`
	ApiGroupName string `json:"apigroupname"`
	Version      string `json:"version"` // 分组版本

}
type CreateApiGroupRes struct {
}

// 删除API
type DeleteApiReq struct {
	g.Meta `path:"deleteapi" method:"delete" tags:"menu" dc:"删除API"`
}

type DeleteApiRes struct {
}

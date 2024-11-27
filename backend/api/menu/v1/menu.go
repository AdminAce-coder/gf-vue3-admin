package v1

import (
	"gf-vue3-admin/internal/service/register"

	"github.com/gogf/gf/v2/frame/g"
)

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
	g.Meta `path:"/capigp" method:"post" tags:"menu" dc:"新增分组"`
	//ApiGroupName string `json:"apigroupname"`
	//Version      string `json:"version"` // 分组版本
	ApiPath  string                    `json:"apipath"`
	Register register.RouteGroupConfig `json:"register"`
}
type CreateApiGroupRes struct {
}

// 删除API
type DeleteapiReq struct {
	g.Meta `path:"/delapi" method:"delete" tags:"menu" dc:"删除API"`
	//apictrl.DeleteApi
	Apipath  string `json:"apipath"`
	ApiGroup string `json:"apigroup"`
	//ApiV string `json:"apiversion"`
}

type DeleteapiRes struct {
}

// // 删除Api组
type DelapigroupReq struct {
	g.Meta    `path:"/delapigp" method:"delete" tags:"menu" dc:"删除API分组"`
	GroupName string `json:"groupname"`
	Version   string `json:"version"`
}

type DelapigroupRes struct {
}

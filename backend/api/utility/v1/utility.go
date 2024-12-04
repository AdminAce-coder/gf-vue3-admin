package v1

import "github.com/gogf/gf/v2/frame/g"

type Utility_Create_OssReq struct {
	g.Meta `path:"/utility_create_oss" method:"post" tags:"utility" dc:"创建存储桶"`
	// 平台
	Platform string `v:"required" dc:"平台"`
	// 存储桶名
	OssName string `v:"required" dc:"存储桶名"`
	// 区域
	Region string `v:"required" dc:"区域"`
}

type Utility_Create_OssRes struct {
}

// 创建SSH连接
type SshInfoReq struct {
	g.Meta `path:"/sshinfo" method:"post" tags:"utility" dc:"创建SSH连接"`
	// 用户
	User string `v:"required" dc:"用户"`
	// 密码
	Password string `v:"required" dc:"密码"`
	Addr     string `v:"required" dc:""`
	Port     int    `v:"required" dc:""`
}

type SshInfoRes struct {
}

// 主机用户
type SshUserReq struct {
	g.Meta `path:"/sshuser" method:"post" tags:"utility" dc:""`
	// 用户
	Hostname string `v:"required" dc:"连接名"`
	User     string `v:"required" dc:"用户"`
	Port     string `v:"required" dc:""`
	Password string `v:"required" dc:""`
	Host     string `v:"required" dc:"地址"`
}

type SshUserRes struct {
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// User is the golang structure of table user for DAO operations like Where/Data.
type User struct {
	g.Meta        `orm:"table:user, do:true"`
	Id            interface{} //
	SelectAccount interface{} // 用户组
	Username      interface{} // 用户名
	Password      interface{} // MD5加密密码
	Nickname      interface{} // 昵称
	Roles         interface{} // 权限集
	Realname      interface{} // 角色名
	Codes         interface{} // 权限码
}

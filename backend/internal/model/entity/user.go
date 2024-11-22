// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// User is the golang structure for table user.
type User struct {
	Id            int    `json:"id"            orm:"id"            description:""`
	SelectAccount string `json:"selectAccount" orm:"selectAccount" description:"用户组"`
	Username      string `json:"username"      orm:"username"      description:"用户名"`
	Password      string `json:"password"      orm:"password"      description:"MD5加密密码"`
	Nickname      string `json:"nickname"      orm:"Nickname"      description:"昵称"`
	Roles         string `json:"roles"         orm:"roles"         description:"权限集"`
	Realname      string `json:"realname"      orm:"realname"      description:"角色名"`
	Codes         string `json:"codes"         orm:"codes"         description:"权限码"`
}

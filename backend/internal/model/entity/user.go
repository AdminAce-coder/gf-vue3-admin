// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// User is the golang structure for table user.
type User struct {
	UserId        int    `json:"userId"        orm:"userId"        description:""`
	SelectAccount string `json:"selectAccount" orm:"selectAccount" description:"用户组"`
	UserName      string `json:"userName"      orm:"userName"      description:"用户名"`
	Password      string `json:"password"      orm:"password"      description:"MD5加密密码"`
	Nickname      string `json:"nickname"      orm:"Nickname"      description:"昵称"`
	Roles         string `json:"roles"         orm:"roles"         description:"权限集"`
	Realname      string `json:"realname"      orm:"realname"      description:"角色名"`
	Codes         string `json:"codes"         orm:"codes"         description:"权限码"`
	Phone         string `json:"phone"         orm:"phone"         description:"手机号"`
	Email         string `json:"email"         orm:"email"         description:"邮箱"`
	Status        string `json:"status"        orm:"status"        description:"状态"`
	HeadPortrait  string `json:"headPortrait"  orm:"head_portrait" description:"头像 BASE64编码"`
	Buttons       string `json:"buttons"       orm:"buttons"       description:""`
}

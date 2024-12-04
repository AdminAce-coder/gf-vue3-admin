// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Ssh is the golang structure of table ssh for DAO operations like Where/Data.
type Ssh struct {
	g.Meta   `orm:"table:ssh, do:true"`
	HostName interface{} //
	User     interface{} //
	Password interface{} //
	Port     interface{} //
	Host     interface{} //
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Ssh is the golang structure for table ssh.
type Ssh struct {
	HostName string `json:"hostName" orm:"HostName" description:""`
	User     string `json:"user"     orm:"User"     description:""`
	Password string `json:"password" orm:"Password" description:""`
	Port     int    `json:"port"     orm:"Port"     description:""`
	Addr     string `json:"addr"     orm:"Addr"     description:""`
}

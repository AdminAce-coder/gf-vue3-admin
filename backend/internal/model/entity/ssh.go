// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Ssh is the golang structure for table ssh.
type Ssh struct {
	HostName string `json:"hostName" orm:"Host_Name" description:""`
	User     string `json:"user"     orm:"User"      description:""`
	Password string `json:"password" orm:"Password"  description:""`
	Port     int    `json:"port"     orm:"Port"      description:""`
	Host     string `json:"host"     orm:"Host"      description:""`
}

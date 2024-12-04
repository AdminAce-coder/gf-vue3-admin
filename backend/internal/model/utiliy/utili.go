package utiliy

var SshUser = &SshUserInfo{}

type SshUserInfo struct {
	HostName string
	User     string
	Password string
	Addr     string
	Port     int
}

// 查询SSH连接信息
type SshConnectInfoInput struct {
	Host string
}

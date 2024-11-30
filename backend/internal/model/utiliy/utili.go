package utiliy

var SshUser = &SshUserInfo{}

type SshUserInfo struct {
	User     string
	Password string
	Addr     string
	Port     int
}

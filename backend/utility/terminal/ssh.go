package terminal

import (
	"bytes"
	"context"
	"gf-vue3-admin/internal/model/utiliy"
	"io"
	"strings"
	"sync"
	"time"

	"github.com/gogf/gf/v2/os/glog"
	gossh "golang.org/x/crypto/ssh"
)

type Sshconfig struct {
	Userinfo    *utiliy.SshUserInfo
	DialTimeOut time.Duration  `json:"dialTimeOut"`
	Client      *gossh.Client  `json:"client"`
	Session     *gossh.Session `json:"session"`
	LastResult  string         `json:"lastResult"`
}

func (s *Sshconfig) NewSshConfig(ctx context.Context) (*Sshconfig, error) {
	glog.Infof(ctx, "Userinfo是,%s", s.Userinfo)
	// 配置超时时间
	if s.DialTimeOut == 0 {
		s.DialTimeOut = time.Second * 5
	}
	config := &gossh.ClientConfig{}
	config.SetDefaults()
	config.User = s.Userinfo.User
	// 配置认证方式为密码认证
	config.Auth = []gossh.AuthMethod{gossh.Password(s.Userinfo.Password)}
	config.HostKeyCallback = gossh.InsecureIgnoreHostKey() // 忽略主机密钥
	proto := "tcp"
	if strings.Contains(s.Userinfo.Addr, ":") {
		proto = "tcp6"
	}
	client, err := gossh.Dial(proto, s.Userinfo.Addr, config) // 连接到 SSH 服务器
	if nil != err {
		return s, err
	}
	if s.Client == nil {
		glog.Infof(ctx, "client 为空")
	}
	s.Client = client
	return s, nil
}

type SshConn struct {
	StdinPipe   io.WriteCloser  // 标准输入管道
	ComboOutput *wsBufferWriter // 组合输出
	Session     *gossh.Session  // SSH 会话
}

// NewSshConn 创建 SSH 连接
func (s *Sshconfig) NewSshConn(cols, rows int) (*SshConn, error) {
	sshSession, err := s.Client.NewSession() // 创建会话
	if err != nil {
		return nil, err
	}

	stdinP, err := sshSession.StdinPipe() // 获取标准输入管道
	if err != nil {
		return nil, err
	}

	comboWriter := new(wsBufferWriter)
	sshSession.Stdout = comboWriter
	sshSession.Stderr = comboWriter

	modes := gossh.TerminalModes{
		gossh.ECHO:          1,     // 回显
		gossh.TTY_OP_ISPEED: 14400, // 输入速度
		gossh.TTY_OP_OSPEED: 14400, // 输出速度
	}
	// 请求 PTY
	if err := sshSession.RequestPty("xterm", rows, cols, modes); err != nil {
		return nil, err
	}
	// 启动 shell
	if err := sshSession.Shell(); err != nil {
		return nil, err
	}
	return &SshConn{StdinPipe: stdinP, ComboOutput: comboWriter, Session: sshSession}, nil
}

func (s *SshConn) Close() {
	if s.Session != nil {
		s.Session.Close()
	}
}

// wsBufferWriter 是组合输出缓冲区的结构体
type wsBufferWriter struct {
	Buffer bytes.Buffer // 缓冲区
	mu     sync.Mutex
}

func (w *wsBufferWriter) Write(p []byte) (int, error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.Buffer.Write(p)
}

func (w *wsBufferWriter) Bytes() []byte {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.Buffer.Bytes()
}

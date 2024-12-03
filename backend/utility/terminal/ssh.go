package terminal

import (
	"bytes"
	"context"
	"fmt"
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
	glog.Infof(ctx, "开始SSH配置，用户信息: %+v", s.Userinfo)

	if s.Userinfo == nil {
		return nil, fmt.Errorf("用户信息未初始化")
	}

	// 确保地址中包含端口号
	if !strings.Contains(s.Userinfo.Addr, ":") {
		s.Userinfo.Addr = fmt.Sprintf("%s:%d", s.Userinfo.Addr, 22)
	}

	config := &gossh.ClientConfig{}
	config.SetDefaults()
	config.User = s.Userinfo.User
	config.Auth = []gossh.AuthMethod{gossh.Password(s.Userinfo.Password)}
	config.HostKeyCallback = gossh.InsecureIgnoreHostKey()

	// 始终使用tcp协议，因为我们使用的是IPv4地址
	proto := "tcp"

	glog.Infof(ctx, "尝试SSH连接到 %s，使用协议：%s", s.Userinfo.Addr, proto)
	client, err := gossh.Dial(proto, s.Userinfo.Addr, config)
	if err != nil {
		return nil, fmt.Errorf("SSH连接失败: %v", err)
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
		gossh.ECHO:          0,     // 禁用回显以防止重复输出
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

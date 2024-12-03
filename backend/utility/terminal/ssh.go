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
	sshSession, err := s.Client.NewSession()
	if err != nil {
		return nil, err
	}

	stdinP, err := sshSession.StdinPipe()
	if err != nil {
		return nil, err
	}

	comboWriter := new(wsBufferWriter)
	sshSession.Stdout = comboWriter
	sshSession.Stderr = comboWriter

	// 修改终端模式设置
	modes := gossh.TerminalModes{
		gossh.ECHO:          1,     // 启用回显
		gossh.TTY_OP_ISPEED: 14400, // 输入速度
		gossh.TTY_OP_OSPEED: 14400, // 输出速度
		gossh.ICANON:        0,     // 禁用规范模式
		gossh.ISIG:          1,     // 启用信号处理
		gossh.ICRNL:         0,     // 禁用将CR转换为NL
		gossh.IEXTEN:        0,     // 禁用扩展处理
		gossh.OPOST:         1,     // 启用输出处理
		gossh.ONLCR:         0,     // 禁用将NL转换为CRNL
		gossh.IXON:          0,     // 禁用XON/XOFF流控
		gossh.IXOFF:         0,     // 禁用XON/XOFF流控
		gossh.IXANY:         0,     // 禁用任意字符重启输出
		gossh.IGNPAR:        0,     // 禁用忽略奇偶校验错误
		gossh.PARMRK:        0,     // 禁用标记奇偶校验错误
		gossh.INPCK:         0,     // 禁用输入奇偶校验
		gossh.ISTRIP:        0,     // 禁用剥离第8位
		gossh.INLCR:         0,     // 禁用将NL转换为CR
		gossh.IGNCR:         0,     // 禁用忽略CR
	}

	// 请求伪终端
	if err := sshSession.RequestPty("xterm-256color", rows, cols, modes); err != nil {
		sshSession.Close()
		return nil, fmt.Errorf("请求PTY失败: %v", err)
	}

	// 启动shell
	if err := sshSession.Shell(); err != nil {
		sshSession.Close()
		return nil, fmt.Errorf("启动Shell失败: %v", err)
	}

	return &SshConn{
		StdinPipe:   stdinP,
		ComboOutput: comboWriter,
		Session:     sshSession,
	}, nil
}

func (s *SshConn) Close() {
	if s.Session != nil {
		s.Session.Close()
	}
}

// wsBufferWriter 是组合输出缓冲区的结构体
type wsBufferWriter struct {
	Buffer bytes.Buffer
	mu     sync.Mutex
}

func (w *wsBufferWriter) Write(p []byte) (int, error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	// 添加调试日志
	fmt.Printf("收到SSH输出: %q\n", string(p))
	return w.Buffer.Write(p)
}

func (w *wsBufferWriter) Bytes() []byte {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.Buffer.Bytes()
}

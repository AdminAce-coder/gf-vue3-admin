package terminal

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"

	"github.com/gorilla/websocket"
)

// WsMsg 是 WebSocket 消息的结构体
type WsMsg struct {
	Type string `json:"type"`
	Data string `json:"data,omitempty"` // WsMsgCmd
	Cols int    `json:"cols,omitempty"` // WsMsgResize
	Rows int    `json:"rows,omitempty"` // WsMsgResize
	Key  string `json:"key,omitempty"`  // 键盘事件
	//Timestamp int `json:"timestamp,omitempty"` // WsMsgHeartbeat
}

// LogicSshWsSession 是逻辑 SSH WebSocket 会话的结构体
type SshWsSession struct {
	//stdinPipe io.WriteCloser  // 标准输入管道
	//session *ssh.Session    // SSH 会话特
	wsConn *websocket.Conn // WebSocket 连接
	sshcon *SshConn
}

func NewSshWsSession(sshcon *SshConn, wsConn *websocket.Conn) *SshWsSession {
	return &SshWsSession{sshcon: sshcon, wsConn: wsConn}
}

// 关闭WSeision
func (s *SshWsSession) Close() {
	s.sshcon.Close()
}

func (s *SshWsSession) Start(quitchan chan bool) {
	go s.receiveWsMsg(quitchan)
	go s.SendComboOutput(quitchan)
}

// 接收WebSocket 消息

func (s *SshWsSession) receiveWsMsg(exitCh chan bool) {
	ctx := gctx.New()
	wscon := s.wsConn
	for {
		select {
		case <-exitCh:
			return
		default:
			_, data, err := wscon.ReadMessage()
			if err != nil {
				glog.Error(ctx, "读取WebSocket消息错误:", err)
				return
			}

			rmsg := WsMsg{}
			if err := json.Unmarshal(data, &rmsg); err != nil {
				glog.Error(ctx, "解析消息失败:", err)
				continue
			}

			// 处理特殊键
			if rmsg.Type == "key" {
				switch rmsg.Key {
				case "ctrl+c":
					// 发送中断信号 (ASCII 0x03 = Ctrl+C)
					if err := s.SendmgsToPipe([]byte{0x03}); err != nil {
						glog.Error(ctx, "发送中断信号失败:", err)
					}
					continue
				}
			}

			// 过滤空命令
			command := strings.TrimSpace(rmsg.Data)
			if command == "" {
				continue
			}

			glog.Infof(ctx, "发送的命令是：%s", command)
			// 将命令转换为字节切片，确保命令以换行符结束
			byteMsg := []byte(command + "\n")
			if err := s.SendmgsToPipe(byteMsg); err != nil {
				glog.Error(ctx, "发送命令到SSH管道失败:", err)
				continue
			}
		}
	}
}

// SendmgsToPipe 发送消息到管道
func (s *SshWsSession) SendmgsToPipe(msg []byte) error {
	if s.sshcon == nil || s.sshcon.StdinPipe == nil {
		return fmt.Errorf("SSH连接未建立")
	}
	_, err := s.sshcon.StdinPipe.Write(msg)
	return err
}

// SendComboOutput 发送组合输出
func (s *SshWsSession) SendComboOutput(exitCh chan bool) {
	ctx := gctx.New()
	wscon := s.wsConn
	// 创建一个定时器，减少检查间隔以提高响应速度
	ticker := time.NewTicker(time.Millisecond * time.Duration(30))
	defer ticker.Stop()

	for {
		select {
		case <-exitCh:
			return
		case <-ticker.C:
			// 从缓冲区读取数据
			input := s.sshcon.ComboOutput.Bytes()
			if len(input) > 0 {
				// 清空缓冲区
				s.sshcon.ComboOutput.Buffer.Reset()

				// 发送输出到WebSocket
				if err := wscon.WriteJSON(WsMsg{
					Type: "cmd",
					Data: string(input),
				}); err != nil {
					glog.Error(ctx, "发送WebSocket消息失败:", err)
					return
				}
			}
		}
	}
}

package terminal

import (
	"encoding/json"
	"fmt"
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

			if err := s.receiveWsMsg(data); err != nil {
				glog.Error(ctx, "处理WebSocket消息失败:", err)
				continue
			}
		}
	}
}

func (s *SshWsSession) receiveWsMsg(msg []byte) error {
	var wsMsg WsMsg
	if err := json.Unmarshal(msg, &wsMsg); err != nil {
		return err
	}

	switch wsMsg.Type {
	case "message":
		// 添加回车换行符
		data := wsMsg.Data + "\r\n"
		if _, err := s.sshcon.Write([]byte(data)); err != nil {
			return fmt.Errorf("写入SSH失败: %v", err)
		}
	case "key":
		if wsMsg.Key == "ctrl+c" {
			if _, err := s.sshcon.Write([]byte{3}); err != nil {
				return fmt.Errorf("发送Ctrl+C失败: %v", err)
			}
		}
	}
	return nil
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
	ticker := time.NewTicker(time.Millisecond * 10) // 减少检查间隔到10ms
	defer ticker.Stop()

	for {
		select {
		case <-exitCh:
			return
		case <-ticker.C:
			input := s.sshcon.ComboOutput.Bytes()
			if len(input) > 0 {
				// 发送输出到WebSocket
				err := wscon.WriteJSON(WsMsg{
					Type: "cmd",
					Data: string(input),
				})
				if err != nil {
					glog.Error(ctx, "发送WebSocket消息失败:", err)
					return
				}
				// 立即清空缓冲区
				s.sshcon.ComboOutput.Reset()
			}
		}
	}
}

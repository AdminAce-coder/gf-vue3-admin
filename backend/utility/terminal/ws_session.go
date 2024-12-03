package terminal

import (
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"time"

	"github.com/gorilla/websocket"
)

// WsMsg 是 WebSocket 消息的结构体
type WsMsg struct {
	Type string `json:"type"`
	Data string `json:"data,omitempty"` // WsMsgCmd
	//Cols      int    `json:"cols,omitempty"`      // WsMsgResize
	//Rows      int    `json:"rows,omitempty"`      // WsMsgResize
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
	// 解析消息
	ctx := gctx.New()
	wscon := s.wsConn
	for {
		select {
		case <-exitCh:
			return
		default:
			_, data, err := wscon.ReadMessage()
			if err != nil {
				fmt.Println(err)
				return
			}
			glog.Info(ctx, "正在解析消息")
			rmgs := WsMsg{}
			if err := json.Unmarshal(data, &rmgs); err != nil {
				return
			}
			glog.Infof(ctx, "命令是%s", rmgs.Data)
			// 将字符串转换为字节切片
			byteMsg := []byte(rmgs.Data)
			s.SendmgsToPipe(byteMsg)
			return
		}

	}

}

// 发送 WebSocket 输入命令到 SSH 会话的标准输入管道
func (s *SshWsSession) SendmgsToPipe(cmdBytes []byte) error {
	_, err := s.sshcon.StdinPipe.Write(cmdBytes)
	return err
}

// 发送 WebSocket 消息
func (s *SshWsSession) SendComboOutput(exitCh chan bool) {
	wscon := s.wsConn
	// 创建一个定时器
	ticker := time.NewTicker(time.Millisecond * time.Duration(60))
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			if s.sshcon.ComboOutput == nil {
				return
			}
			// 获取输入
			input := s.sshcon.ComboOutput.Bytes()
			if len(input) > 0 {
				wsdata, err := json.Marshal(WsMsg{
					Type: "cmd",
					Data: string(input),
				})
				if err != nil {
					fmt.Println("获取数据失败", err.Error())
					continue
				}
				// 将组合输出发送回 WebSocket
				err = wscon.WriteMessage(websocket.TextMessage, wsdata)
				if err != nil {
					fmt.Println("发送数据失败", err.Error())
				}
				s.sshcon.ComboOutput.Buffer.Reset() // 重置组合输出缓冲区
			}
		case <-exitCh:
			return
		}
	}
}

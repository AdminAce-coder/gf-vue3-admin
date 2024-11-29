package cmd

import (
	"fmt"
	"net/http"
	"sync/atomic"

	"github.com/olahol/melody"
)

var idCounter atomic.Int64 // 用于生成唯一的ID

func SrartWebsoket() {
	m := melody.New()

	// 配置消息大小限制
	m.Config.MaxMessageSize = 1024 * 1024 * 10 // 10MB
	m.Config.MessageBufferSize = 1024 * 1024   // 1MB

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		m.HandleRequest(w, r)
	})

	m.HandleConnect(func(s *melody.Session) {
		id := idCounter.Add(1) //这是

		s.Set("id", id)

		s.Write([]byte(fmt.Sprintf("iam %d", id)))
	})

	m.HandleDisconnect(func(s *melody.Session) {
		if id, ok := s.Get("id"); ok {
			m.BroadcastOthers([]byte(fmt.Sprintf("dis %d", id)), s)
		}
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		if id, ok := s.Get("id"); ok {
			m.BroadcastOthers([]byte(fmt.Sprintf("set %d %s", id, msg)), s)
		}
	})

	http.ListenAndServe(":6000", nil)
	// 启动
	fmt.Printf("已启动websocker服务端..6000")
}

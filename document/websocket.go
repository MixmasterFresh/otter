package document

import (
	"encoding/json"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait  = 15 * time.Second
	pingPeriod = 10 * time.Second
	pongWait   = 15 * time.Second
)

func (user *User) read() {
	conn := user.conn
	defer conn.Close()
	conn.SetReadLimit(1024)
	conn.SetReadDeadline(time.Now().Add(pongWait))
	conn.SetPongHandler(func(string) error {
		conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		if len(msg) != 0 {
			user.handleMessage(msg)
		}
	}
}

func (user *User) write() {
	conn := user.conn
	pingTicker := time.NewTicker(pingPeriod)
	defer pingTicker.Stop()
	defer conn.Close()
	for {
		select {
		case msg := <-user.sendQueue:
			if len(msg) != 0 {
				conn.SetWriteDeadline(time.Now().Add(writeWait))
				if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
					return
				}
			}
		case <-pingTicker.C:
			conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

type message struct {
	Event   string          `json:"event"`
	EventId string          `json:"eventId"`
	Data    json.RawMessage `json:"data"`
}

type insertion struct {
	Id      string `json:"id"`
	Index   int    `json:"index"`
	Content string `json:"content"`
	Parent  string `json:"parent"`
}

type deletion struct {
	Id     string `json:"id"`
	Index  int    `json:"index"`
	Length int    `json:"length"`
	Parent string `json:"parent"`
}

func (user *User) handleMessage(msg []byte) {
	var decodedMsg message
	err := json.Unmarshal(msg, &decodedMsg)
	if err != nil {
		return
	}
	user.document.broadcastToAllExcept(msg, user.id)
	switch decodedMsg.Event {
	case "insert":
		var insertMsg insertion
		err = json.Unmarshal(decodedMsg.Data, &insertMsg)
		user.queueInsertion(insertMsg, msg)
	case "delete":
		var deleteMsg deletion
		err = json.Unmarshal(decodedMsg.Data, &deleteMsg)
		user.queueDeletion(deleteMsg, msg)
	}
}

func (user *User) queueInsertion(msg insertion, rawMsg []byte) {
	
}

func (user *User) queueDeletion(msg deletion, rawMsg []byte) {

}

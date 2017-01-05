package document

import (
	"encoding/json"
	"time"
	"github.com/TheAustinSeven/otter/operation"
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
	User	string			`json:"user"`
	Data    json.RawMessage `json:"data"`
}

type insertion struct {
	Id        int   			`json:"id"`
	Index     int      			`json:"index"`
	Content   string   			`json:"content"`
	Parent    string   			`json:"parent"`
	Ancestors map[string][]int `json:"ancestors"`
}

type deletion struct {
	Id        int   `json:"id"`
	Index     int      `json:"index"`
	Length    int      `json:"length"`
	Parent    string   `json:"parent"`
	Ancestors map[string][]int `json:"ancestors"`
}

type notifier struct {
	Id        int   `json:"id"`
	Index     int      `json:"index"`
	Ancestors map[string][]int `json:"ancestors"`
}

func (user *User) handleMessage(msg []byte) {
	var decodedMsg message
	err := json.Unmarshal(msg, &decodedMsg)
	if err != nil {
		return
	}else if decodedMsg.User != user.id {
		return
	}else{
		//rebroadcast
	}
	switch decodedMsg.Event {
	case "insert":
		var insertMsg insertion
		err = json.Unmarshal(decodedMsg.Data, &insertMsg)
		if err != nil {
			//log an error and write an error back
		}
		user.queueInsertion(insertMsg, msg)
	case "delete":
		var deleteMsg deletion
		err = json.Unmarshal(decodedMsg.Data, &deleteMsg)
		if err != nil {
			//log an error and write an error back
		}
		user.queueDeletion(deleteMsg, msg)
	case "notifier":
		var notifierMsg notifier
		err = json.Unmarshal(decodedMsg.Data, &notifierMsg)	
		if err != nil {
			//log an error and write an error back
		}
		user.queueNotifier(notifierMsg, msg)
	}
}

func (user *User) queueInsertion(msg insertion, rawMsg []byte) {
	op := operation.Operation{
		Id: msg.Id,
		Author: user.id,
		Format: operation.INSERT,
		Index: msg.Index,
		Length: len(msg.Content),
		Contents: msg.Content,
		Ancestors: msg.Ancestors,
	}

	//do something with ancestors to close the trees if possible. Broadcast to all except...
}

func (user *User) queueDeletion(msg deletion, rawMsg []byte) {
	op := operation.Operation{
		Id: msg.Id,
		Author: user.id,
		Format: operation.INSERT,
		Index: msg.Index,
		Length: msg.Length,
		Ancestors: msg.Ancestors,
	}
}

func (user *User) queueNotifier(msg notifier, rawMsg []byte) {
	op := operation.Operation{
		Id: msg.Id,
		Author: user.id,
		Format: operation.NOTIFIER,
		Index: msg.Index,
		Length: 0,
		Ancestors: msg.Ancestors,
	}
}

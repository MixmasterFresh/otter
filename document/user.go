package document

import (
	"fmt"
	"github.com/TheAustinSeven/otter/auth"
	"github.com/gorilla/websocket"
	"net/http"
)

// User represents the individual user as they pertain to documents
type User struct {
	id       string
	key      string
	hasKey   bool
	document *Document
	conn     *websocket.Conn
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// NewUser creates a new user
func NewUser(id string, document *Document) *User {
	user := new(User)
	user.id = id
	user.key = auth.GenerateKey(32)
	user.document = document
	return user
}

// OpenConnection opens a websocket connection with the user
func (user *User) OpenConnection(w http.ResponseWriter, r *http.Request) {
	if !user.hasKey {
		return
	}
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Failed to set websocket upgrade: %s\n", err.Error())
		return
	}
	if user.authenticate(conn) {
		user.conn = conn
		user.hasKey = false
	}
}

func (user *User) authenticate(conn *websocket.Conn) bool {
	_, proposedKey, _ := conn.ReadMessage()
	return string(proposedKey) == user.key
}

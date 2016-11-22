package document

import (
	"fmt"
	"net/http"

	"github.com/TheAustinSeven/otter/auth"
	"github.com/gorilla/websocket"
)

// User represents the individual user as they pertain to documents
type User struct {
	id        string
	key       string
	hasKey    bool
	document  *Document
	conn      *websocket.Conn
	sendQueue chan []byte
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
	user.sendQueue = make(chan []byte, 100)
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

	user.conn = conn
	go user.read()
	go user.write()
}

func (user *User) ResetKey() string {
	user.key = auth.GenerateKey(32)
	return user.key
}

func (user *User) Authenticate(proposedKey string) bool {
	return proposedKey == user.key
}

func (user *User) CreateCookie() *http.Cookie {

}

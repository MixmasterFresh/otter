package document

import (
	"net/http"
    "github.com/gorilla/websocket"
)

// User represents the individual user as they pertain to documents
type User struct {
	id   string
	document *Document
	conn *websocket.Conn
}

// NewUser creates a new user
func NewUser(id string, document *Document) *User {
	user := new(User)
	user.id = id
	user.document = document
	return user
}

// OpenConnection opens a websocket connection with the user
func (user *User) OpenConnection(w http.ResponseWriter, r *http.Request) {
}
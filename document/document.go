package document

import (
	"strconv"
	"sync"

	"github.com/TheAustinSeven/otter/document/rope"
	"github.com/TheAustinSeven/otter/operation"
)

// Document represents the individual documents uploaded to Otter
type Document struct {
	users        map[string]*User
	messages     chan *operation.Operation
	operations   chan *operation.Operation
	contents     *rope.Rope
	inEditedList bool
	inEmptyList  bool
	editsList    map[string]*Document
	emptyList    map[string]*Document
	mutex        sync.Mutex
}

// NewDocument returns a document with the starting string and otherwise empty values
func NewDocument(contents string, editsList map[string]*Document, emptyList map[string]*Document) *Document {
	document := new(Document)
	document.contents = rope.NewRope(contents)
	document.inEditedList = false
	document.inEmptyList = false
	document.users = make(map[string]*User)
	document.editsList = editsList
	document.emptyList = emptyList
	return document
}

// GetString returns contents as a string
func (document *Document) GetString() string {
	return document.contents.ToString()
}

// GetMetadata returns a map of metadata
func (document *Document) GetMetadata() map[string]string {
	data := make(map[string]string)
	length := document.contents.Size()
	data["length"] = strconv.Itoa(length)
	userCount := len(document.users)
	data["users"] = strconv.Itoa(userCount)
	return data
}

// AddUser facilitates the adding of users to a document
func (document *Document) AddUser() {
	if len(document.users) == 0 {
		document.operations = make(chan *operation.Operation, 16)
		go document.listenForOperations()
	}
}

// RemoveUser facilitates the removal of users from a document
func (document *Document) RemoveUser() {
	if len(document.users) == 0 {
		close(document.operations)
	}
}

func (document *Document) resizeChannels() {

}

func (document *Document) enqueOperation(op *operation.Operation) {
	document.mutex.Lock()
	defer document.mutex.Unlock()
	//transform operation
}

func (document *Document) listenForOperations() {
	for working := true; working; {
		op, working := <-document.operations
		if working {
			//transform
			document.enqueOperation(op)
		}
	}
	//Safely handle shutdown
}

func (document *Document) broadcastToAllExcept(msg []byte, id string) {
	for userID, user := range document.users {
		if userID != id {
			user.sendQueue <- msg
		}
	}
}

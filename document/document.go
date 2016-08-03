package document

import (
	"strconv"
	"otter/document/rope"
	"otter/operation"
	"otter/user"
)

// Document represents the individual documents uploaded to Otter
type Document struct {
	users          map[string]*user.User
	operations     chan *operation.Operation
	contents       *rope.Rope
	InEditedList   bool
}

// NewDocument returns a document with the starting string and otherwise empty values
func NewDocument(contents string) *Document{
	document := new(Document)
	document.contents = rope.NewRope(contents)
	document.operations = make(chan *operation.Operation, 64)
	document.InEditedList = false;
	document.users = make(map[string]*user.User)
	return document
}

func (document *Document) getString() string{
	return document.contents.ToString()
}

func (document *Document) getMetadata() map[string]string{
	data := make(map[string]string)
	length := document.contents.Size()
	data["length"] = strconv.Itoa(length)
	userCount := len(document.users)
	data["users"] = strconv.Itoa(userCount)
	return data
}
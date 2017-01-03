package operation

import ()

type OperationCollection struct {
	Operations map[string]*Operation
	histories  map[string]*AuthorQueue
}

type AuthorQueue struct {
	userId string
	list *OperationList
}

// Operation represents the operation being performed
type Operation struct {
	Id        string
	Author	  string
	Format    int
	Index     int
	Length    int
	Contents  string
	Ancestors map[string]string
}

const (//Due to the fact that go has... shortcomings, this is the way to represent enums. *shakes fist*
	NOTIFIER = iota
	INSERT = iota
	DELETE = iota
)

func (store *OperationCollection) Push(op *Operation){

} 

func (store *OperationCollection) Next() *Operation {
	//Collect next 
	return nil
}

func (store *OperationCollection) Cleanup() {

}

func (store *OperationCollection) AddOperation() {

}


package operation

import ()

type OperationCollection struct {
	newOperations *operationBuffer
	histories  map[string]*AuthorQueue
	block chan bool
	processedOperations chan *Operation
}

type AuthorQueue struct {
	userId string
	list *OperationList
}

// Operation represents the operation being performed
type Operation struct {
	Id        int
	Author	  string
	Format    int
	Index     int
	Length    int
	Contents  string
	Ancestors map[string][]int
}

const (//Due to the fact that go has... shortcomings, this is the way to represent enums. *shakes fist*
	NOTIFIER = iota
	INSERT = iota
	DELETE = iota
)

func newOperationCollection() *OperationCollection {
	store := &OperationCollection{}
	store.newOperations = &operationBuffer{}
	store.newOperations.init()
	store.histories = make(map[string]*AuthorQueue)
	store.block = make(chan bool)
	store.processedOperations = make(chan *Operation, 100)
	go store.processOperations()
	return store
}

func (store *OperationCollection) Close() {
	// To be implemented later
}

func (store *OperationCollection) Push(op *Operation){
	store.newOperations.push(op)
	select { case <- store.block: } //do nothing. We are just allowing the channel to send again
} 

func (store *OperationCollection) processOperations() {
	var op *Operation
	for {
		for op = store.newOperations.pop(); op == nil; op = store.newOperations.pop() {
			store.block <- true//potentially more complex check to aid in closing
		}
		newOps := store.histories[op.Author].list.transformOverList(*op, op.Ancestors)
		store.addToAllExcept(op.Author, newOps)
		for _,finalOp := range newOps {
			store.processedOperations <- &finalOp
		}
	}
}

func (store *OperationCollection) Pop() *Operation {
	return <-store.processedOperations
}

func (store *OperationCollection) addToAllExcept(author string, ops []Operation) {
	for key, history := range store.histories {
		if key != author {
			for _, op := range ops {
				history.list.insertOperation(op.simplify())
			}
		}
	}
}


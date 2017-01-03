package operation

type OperationList struct {
    head *simpleOperation
    tail *simpleOperation
}

type simpleOperation struct {
    id      string
    author  string
	format 	int
	index	int
	length  int
	next    *simpleOperation
}

func (list OperationList) insertOperation(op *simpleOperation) {
    if list.head == nil {
        list.head = op
        list.tail = op
    }else{
        list.tail.next = op
        list.tail = op
    }  
}

func (list OperationList) removeByIds(ids map[string]bool) {
    
}

func (op Operation) simplify() simpleOperation {

}




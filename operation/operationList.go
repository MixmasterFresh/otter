package operation

type OperationList struct {
    head *simpleOperation
    tail *simpleOperation
}

type simpleOperation struct {
    id      int
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

func (list OperationList) transformOverList(op Operation, ancestors map[string][]int) []Operation{
    if list.head == nil {
        return []Operation{op}
    }else{
        operations := []Operation{op}
        var previous *simpleOperation  
        for current := list.head; current != nil; current = current.next {
            var newOperations []Operation
            if inSlice(op,ancestors[op.Author]){
                list.remove(current,previous)
                
                
            }else{
                for _, newOp := range operations {
                    transformedOps := newOp.transform(current)
                    if transformedOps != nil {
                        newOperations = append(newOperations, transformedOps...)
                    }
                }
                if len(newOperations) == 0 {
                    return nil
                }else{
                    operations = newOperations
                }
            }
            previous = current
        }
        if len(operations) > 0 {
            return operations
        }else{
            return nil
        }
    }
}

func inSlice(op Operation, ids []int) bool {
    for _, id := range ids {
        if op.Id == id {
            return true
        }
    }
    return false
}

func (list OperationList) remove(current *simpleOperation, previous *simpleOperation) {
    if previous == nil {
        list.head = current.next
    }else{
        previous.next = current.next
    }

    if current.next == nil {
        list.tail = previous
    }
} 

func (op Operation) simplify() *simpleOperation {
    return &simpleOperation{
        id: op.Id,
        author: op.Author,
        format: op.Format,
        index: op.Index,
        length: op.Length,
    }
}




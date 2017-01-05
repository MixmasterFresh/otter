package operation

import (
    "testing"
)

func TestEntries(t *testing.T) {
    buf := &operationBuffer{}
	buf.init()
	go insert(buf,0,1000)
	go insert(buf,1000,1000)
	insert(buf,2000,1000)
    insert(buf,3000,1000)

	i := 0
	var op *Operation
	for ; true;i++ {
		op = buf.pop()
		if op == nil {
			break
		}
	}
    if i != 4000 {
        t.Error("Expected 4000, got ", i)
    }
}

func TestBufferDoublePop(t *testing.T) {
	buf := &operationBuffer{}
	buf.init()
	op := buf.pop()
	if op != nil {
		t.Error("Expected nil, got ", op)
	}
	buf.push(&Operation{})
	op = buf.pop()
	if op == nil {
		t.Error("Expected op not to be nil, but op was nil")
	}
	op = buf.pop()
	if op != nil {
		t.Error("Expected nil, got ", op)
	}
}

func insert(buf *operationBuffer, start int, length int){
	for i := start; i < start + length; i++ {
		op := &Operation{Index: i}
		buf.push(op)
	}
}
package operation

import(
    "sync/atomic"
    "unsafe"    
)

type operationBuffer struct {
    head *operationContainer
    tail *operationContainer
    unset *operationContainer
}

type operationContainer struct {
    op *Operation
    next *operationContainer
}

func (buf *operationBuffer) init(){
    buf.unset = &operationContainer{}
    buf.head = buf.unset
    buf.tail = buf.unset
}

func (buf *operationBuffer) push(rawOp *Operation) {
    container := &operationContainer{op: rawOp, next: buf.unset}
    
    if buf.head == buf.unset {
        if atomic.CompareAndSwapPointer(
                (*unsafe.Pointer)(unsafe.Pointer(&buf.head)),
                unsafe.Pointer(buf.unset),
                unsafe.Pointer(container)) {
                    
            atomic.CompareAndSwapPointer(
                (*unsafe.Pointer)(unsafe.Pointer(&buf.tail)),
                (unsafe.Pointer(buf.unset)),
                (unsafe.Pointer(container)))
            return
        }
    }
    old := buf.tail
    for ; old == buf.unset; old = buf.tail {
    }
    for ; !atomic.CompareAndSwapPointer(
            (*unsafe.Pointer)(unsafe.Pointer(&old.next)),
            unsafe.Pointer(buf.unset),
            unsafe.Pointer(container)) ;  {
        old = old.next
    }
    atomic.CompareAndSwapPointer(
        (*unsafe.Pointer)(unsafe.Pointer(&buf.tail)),
        unsafe.Pointer(old),
        unsafe.Pointer(container))
}

func (buf *operationBuffer) pop() *Operation{
    old := buf.head
    if old == buf.unset {
        return nil
    }
    for ; !atomic.CompareAndSwapPointer(
            (*unsafe.Pointer)(unsafe.Pointer(&buf.head)),
            unsafe.Pointer(old),
            unsafe.Pointer(old.next)) ;  {
        if old == buf.unset {
            return nil
        }
        old = buf.head
    }
    if old == buf.unset {
        return nil
    }
    return old.op
}
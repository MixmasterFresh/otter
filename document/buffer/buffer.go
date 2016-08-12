package buffer

import (
    "github.com/TheAustinSeven/otter/operation"
    "sync"
)

// Buffer is ...
type Buffer struct {
    fragments  []*bufferFragment
    currentWrite int // the index of the bufferFragment we are writing to
    currentRead int // the index of the bufferFragment we are reading from
    writeMutex sync.Mutex
    readMutex sync.Mutex
}

type bufferFragment struct {
    index int // smallest index in the bufferFragment
    start int // where that index is
    end int // where the end is(will be negative unless fragment has become obsolete)
    writeIndex int
    readIndex int
    list  []*operation.Operation
}

// NewBuffer creates a new Buffer
func NewBuffer(length int) *Buffer {
    fragment := new(bufferFragment)
    fragment.list = make([]*operation.Operation, length)
    fragment.index = 0
    fragment.start = 0

    buffer := new(Buffer)
    buffer.fragments = make([]*bufferFragment, 1)
    buffer.fragments[0] = fragment
    buffer.currentWrite = 0
    buffer.currentRead = 0

    return buffer
}

// Push is a thread-safe push operation
func (buffer *Buffer) Push(op *operation.Operation) {
    buffer.writeMutex.Lock()
    defer buffer.writeMutex.Unlock()
    buffer.UnsafePush(op)
}

// Pop is a thread-safe pop operation
func (buffer *Buffer) Pop() (op *operation.Operation) {
    buffer.readMutex.Lock()
    defer buffer.readMutex.Unlock()
    return buffer.UnsafePop()
}

// UnsafePush will be faster than Push, but is not thread-safe
func (buffer *Buffer) UnsafePush(op *operation.Operation) {
    fragment := buffer.fragments[buffer.currentWrite]
    fragment.push(op)
    if fragment.emptyOrFull() { // A fragment that was just pushed cannot be empty
        buffer.provisionNewFragment()
    }
}

// UnsafePop will be faster than Pop, but it is not thread-safe
func (buffer *Buffer) UnsafePop() *operation.Operation {
    fragment := buffer.fragments[buffer.currentRead]
    op := fragment.pop()
    if fragment.readIndex == -1 {
        buffer.currentRead++
    }
    return op
}

func (buffer *Buffer) ReadBetween(start int, end int) []*operation.Operation {
    buffer.readMutex.Lock()
    defer buffer.readMutex.Unlock()
    return buffer.UnsafeReadBetween(start, end)
}

func (buffer *Buffer) ReadAt(index int) *operation.Operation {
    buffer.readMutex.Lock()
    defer buffer.readMutex.Unlock()
    return buffer.UnsafeReadAt(index)
}

func (buffer *Buffer) UnsafeReadBetween(start int, end int) []*operation.Operation {
    startFragmentIndex := buffer.findIndex(start)
    endFragmentIndex := buffer.findIndex(end)
    if startFragmentIndex == -1 || endFragmentIndex == -1 {
        return nil
    }
    read := make([]*operation.Operation, end - start)
    var readTo int
    for fragmentIndex, copyIndex := startFragmentIndex, 0; fragmentIndex <= endFragmentIndex; fragmentIndex++ {
        fragment := buffer.fragments[fragmentIndex]
        if end > fragment.index + len(fragment.list) {
            readTo = fragment.index + len(fragment.list)
        }else{
            readTo = end
        }
        fragment.readRangeIntoSlice(start + copyIndex, readTo, copyIndex, read)
        copyIndex = readTo - start
    }
    return read
}

func (buffer *Buffer) UnsafeReadAt(index int) *operation.Operation {
    fragment := buffer.fragments[buffer.findIndex(index)]
    retrievalIndex := wrappingAdd(fragment.start, index - fragment.index, len(fragment.list))
    return fragment.list[retrievalIndex]
}

func (buffer *Buffer) CleanUntil(index int) {
    buffer.readMutex.Lock()
    defer buffer.readMutex.Unlock()
    buffer.writeMutex.Lock()
    defer buffer.writeMutex.Unlock()
    buffer.UnsafeCleanUntil(index)
}

func (buffer *Buffer) UnsafeCleanUntil(index int) {
    fragmentIndex := buffer.findIndex(index)
    fragment := buffer.fragments[buffer.currentRead]
    currentReadingIndex := fragment.index + wrappingDistance(fragment.start, fragment.readIndex, len(fragment.list))
    if fragmentIndex == -1 || index > currentReadingIndex {
        return
    }else if fragmentIndex > 0 {
        fragments := make([]*bufferFragment, len(buffer.fragments) - fragmentIndex)
        copy(fragments, buffer.fragments[fragmentIndex:])
        buffer.fragments = fragments
        buffer.currentRead -= fragmentIndex
        buffer.currentWrite -= fragmentIndex
    }
    fragment = buffer.fragments[0]
    offset := index - fragment.index
    fragment.index = index
    fragment.start = wrappingAdd(fragment.start, offset, len(fragment.list))
}

func (buffer *Buffer) provisionNewFragment() {
    fragment := new(bufferFragment)
    fragment.list = make([]*operation.Operation, len(buffer.fragments[buffer.currentWrite].list) * 2)
    fragment.index = 0
    fragment.start = 0
    buffer.fragments = append(buffer.fragments, fragment)
}

func (buffer *Buffer) findIndex(index int) int {
    for index, fragment := range buffer.fragments { // Looping through is practical since we can have at most 64 elements in the slice
        if index >= fragment.index {
            if index < fragment.index + wrappingDistance(fragment.start, fragment.writeIndex, len(fragment.list)) {
                return index
            }
        }
    }
    return -1
}

// Fragment Operations

func (fragment *bufferFragment) push(op *operation.Operation) {
    fragment.list[fragment.writeIndex] = op
    fragment.incrementWriteIndex()
}

func (fragment *bufferFragment) pop() *operation.Operation {
    if fragment.readIndex == fragment.writeIndex {
        return nil
    }
    op := fragment.list[fragment.readIndex]
    fragment.incrementReadIndex()
    return op
}

func (fragment *bufferFragment) incrementWriteIndex() {
    fragment.writeIndex++
    if fragment.writeIndex == len(fragment.list) {
        fragment.writeIndex = 0
    }
}

func (fragment *bufferFragment) incrementReadIndex() {
    if fragment.readIndex == fragment.end {
        fragment.readIndex = -1
    }else{
        fragment.readIndex++
        if fragment.readIndex == len(fragment.list) {
            fragment.readIndex = 0
        }
    }
}

func (fragment *bufferFragment) readRangeIntoSlice(start int, end int, offset int, slice []*operation.Operation) {
    length := len(fragment.list)
    distance := end - start
    start -= fragment.index
    end -= fragment.index
    if end > length {
        end -= length
        copy(slice[offset : offset + length - start], fragment.list[start : length])
        copy(slice[offset + length - start : offset + distance], fragment.list[: end])
    }else{
        copy(slice[offset:offset + distance], fragment.list[start : end ])
    }
}

func (fragment *bufferFragment) emptyOrFull() bool {
    return fragment.start == fragment.writeIndex
}

func wrappingDistance(start int, end int, max int) int {
    if end > start {
        return end - start
    }else{
        return (max - start) + end
    }
}

func wrappingAdd(start int, distance int, max int) int {
    total := start + distance
    return total % max
}


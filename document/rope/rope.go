package rope

// Rope is a data structure that represents high volitility strings(changed often) in a tree format.
type Rope struct {
	parent   *Rope
	left     *Rope
	right    *Rope
	contents string //Speed up may be possible by following the 4 byte multiplicity rule and replacing string with array
	split    int    //refers to the dividing index between the left and right node relative to the subtree where this rope is the root
	length   int
}

// MAX_CONTENT_LENGTH is the point where the rope splits a string. Some tweaking may be necessary for peak performance.
const MAX_CONTENT_LENGTH = 32

// NewRope creates a new rope
func NewRope(content string) *Rope {
	store := new(Rope)
	store.contents = content
	store.length = int(len(content))
	return store
}

//Helper Methods

func (rope *Rope) hasLeft() bool {
	return !(rope.left == nil)
}

func (rope *Rope) hasRight() bool {
	return !(rope.right == nil)
}

func (rope *Rope) isEmpty() bool {
	return (rope.length == 0 && rope.contents != "")
}

func (rope *Rope) isLeaf() bool {
	return (rope.left == nil && rope.right == nil)
}

func (rope *Rope) isLeft() bool {
	if rope.parent != nil {
		return rope.parent.left == rope
	}
	return false
}

func (rope *Rope) isRoot() bool {
	return rope.parent == nil
}

// Size returns the length of the rope contents
func (rope *Rope) Size() int {
	return rope.length
}

//Primary Methods

// ToString returns a string representation of the rope
func (rope *Rope) ToString() string {
	if rope.hasLeft() && rope.hasRight() {
		return rope.left.ToString() + rope.right.ToString()
	} else if rope.hasLeft() {
		return rope.left.ToString()
	} else {
		return rope.contents
	}
}

func (rope *Rope) shatterTo(index int) {
	if rope.length > MAX_CONTENT_LENGTH && rope.isLeaf() {
		var leftLength int
		leftLength = rope.length / 2

		rope.splitAt(leftLength)

		//Go deeper
		if index >= leftLength {
			rope.right.shatterTo(index - leftLength)
		} else {
			rope.left.shatterTo(index)
		}

	} else if !rope.isLeaf() {
		if index >= rope.split {
			rope.right.shatterTo(index - rope.split)
		} else {
			rope.left.shatterTo(index)
		}
	}
}

// Insert adds a string at the given index
func (rope *Rope) Insert(index int, content string) bool {
	//Split up rope if necessary
	if rope.length > MAX_CONTENT_LENGTH && rope.isLeaf() {
		rope.shatterTo(index)
	}

	if rope.length < index {
		return false
	}

	length := int(len(content))

	//Perform the insert
	if !rope.isLeaf() {
		if index > rope.split {
			rope.right.Insert(index-rope.split, content)

		} else if index < rope.split {
			rope.split += length
			rope.left.Insert(index, content)

		} else if rope.left.length > rope.right.length {
			rope.pushOutRight()
			newNode := &Rope{parent: rope.right, length: length, contents: content}
			rope.right.left = newNode

			rope.right.split += length
			rope.right.length += length

		} else {
			rope.pushOutLeft()
			newNode := &Rope{parent: rope.left, length: length, contents: content}
			rope.left.right = newNode

			rope.split += length
			rope.left.length += length

		}
	} else {
		if length+rope.length > MAX_CONTENT_LENGTH {
			rope.splitAt(index)
			rope.Insert(index, content)
			rope.length -= length

		} else {
			part1 := rope.contents[0:index]
			part2 := rope.contents[index:rope.length]
			rope.contents = part1 + content + part2

		}
	}
	rope.length += length
	return true
}

// Delete removes a string of a given length from the given index
func (rope *Rope) Delete(index int, length int) bool {
	if length == 0 {
		return true
	}

	if index+length > rope.length {
		return false
	}

	if rope.isLeaf() {
		rope.removeContent(index, length)
	} else {
		maxLeftLength := rope.split - index
		var leftLength int
		if maxLeftLength > length {
			leftLength = length
		} else if maxLeftLength < 0 {
			leftLength = 0
		} else {
			leftLength = maxLeftLength
		}
		rightLength := length - leftLength

		rightInd := index - rope.split
		if rightInd < 0 {
			rightInd = 0
		}

		if leftLength > 0 {
			if leftLength == rope.left.length {
				rope.left.destroy()

			} else {
				rope.left.Delete(index, leftLength)

			}
			rope.split -= leftLength
		}

		if rightLength > 0 {
			if rightLength == rope.right.length {
				rope.right.destroy()

			} else {
				rightIndex := rightInd
				rope.right.Delete(rightIndex, rightLength)

			}
		}
	}
	rope.length -= length

	if !(rope.hasLeft() && rope.hasRight()) && !(rope.isLeaf()) {
		rope.truncate()
	}

	return true
}

func (rope *Rope) removeContent(index int, length int) {
	if index == 0 {
		rope.contents = rope.contents[length:rope.length]
	} else if index+length >= rope.length {
		rope.contents = rope.contents[0:index]
	} else {
		rope.contents = rope.contents[0:index] + rope.contents[index+length:rope.length]
	}
}

func (rope *Rope) pushOutLeft() {
	newLeft := &Rope{parent: rope, length: rope.left.length, split: rope.left.length, left: rope.left}
	oldLeft := rope.left
	oldLeft.parent = newLeft
	rope.left = newLeft
}

func (rope *Rope) pushOutRight() {
	newRight := &Rope{parent: rope, length: rope.right.length, split: 0, right: rope.right}
	oldRight := rope.right
	oldRight.parent = newRight
	rope.right = newRight
}

func (rope *Rope) splitAt(index int) {
	//Create new ropes
	newLeft := &Rope{parent: rope, length: index}
	newRight := &Rope{parent: rope, length: rope.length - index}

	newLeft.contents = rope.contents[0:index]
	newRight.contents = rope.contents[index:rope.length]

	//Clean up old rope
	rope.contents = "" // may be alright to change to nil, but I think i'll play it safe for now.
	rope.left = newLeft
	rope.right = newRight
	rope.split = index
}

func (rope *Rope) truncate() {
	if rope.isLeaf() {
		rope.destroy()
	} else if rope.isEmpty() {
		if (rope.hasLeft() && rope.hasRight()) || rope.isRoot() {
			return
		}
		parent := rope.parent
		var child *Rope
		if rope.hasLeft() {
			child = rope.left
		} else {
			child = rope.right
		}

		if rope.isLeft() {
			parent.left = child
		} else {
			parent.right = child
		}

		child.parent = parent

		//Clean up afterwards
		rope.left = nil
		rope.right = nil
		rope.parent = nil
	}
}

func (rope *Rope) destroy() {
	if !rope.isLeaf() {
		if rope.hasLeft() {
			rope.left.destroy()
			rope.left = nil
		}
		if rope.hasRight() {
			rope.right.destroy()
			rope.right = nil
		}
	}

	if rope.isLeft() {
		rope.parent.left = nil
	} else {
		rope.parent.right = nil
	}

	rope.parent = nil
}

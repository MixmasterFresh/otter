package operation

import ()

// Operation represents the operation being performed
type Operation struct {
	insertion bool //may want to try and follow the rule of fours here
	index     int
	length    int
	contents  string
	parent    int
}

// Transform applies a transformation
func (first *Operation) Transform(second *Operation) {

}

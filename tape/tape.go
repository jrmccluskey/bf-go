// Package tape contains the data representation for a bf program.
package tape

// DataState encapsulates the tape representation of a bf program's memory.
type DataState struct {
	tape  []byte
	index uint16
}

const (
	tapeLength = 30000
	startIndex = 0
)

// NewDataState initializes a new BF tracking state struct, allocating memory for the tape and a stack for looping.
func NewDataState() *DataState {
	return &DataState{tape: make([]byte, tapeLength, tapeLength), index: startIndex}
}

// IncrementPointer moves the data pointer to the next cell.
// Corresponds to the ">" command.
func (s *DataState) IncrementPointer() {
	s.index++
}

// DecrementPointer moves the data pointer to the previous cell.
// Corresponds to the "<" command.
func (s *DataState) DecrementPointer() {
	s.index--
}

// IncrementByte increases the value of the current cell by one.
// Corresponds to the "+" command.
func (s *DataState) IncrementByte() {
	s.tape[s.index]++
}

// DecrementByte decreases the value of the current cell by one.
// Corresponds to the "-" command.
func (s *DataState) DecrementByte() {
	s.tape[s.index]--
}

// OutputByte returns the value of the byte in the current cell.
// Corresponds to the "." command.
func (s *DataState) OutputByte() byte {
	return s.tape[s.index]
}

// AcceptByte takes a single-byte argument and stores it in the current cell.
// Corresponds to the "," command.
func (s *DataState) AcceptByte(input byte) {
	s.tape[s.index] = input
}

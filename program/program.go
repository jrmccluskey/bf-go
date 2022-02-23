// Package program contains program representation for a bf program.
package program

import (
	"errors"
	"fmt"

	"github.com/emirpasic/gods/stacks/linkedliststack"
)

// InstructionState encapsulates program execution information for a bf program.
type InstructionState struct {
	instructions   []byte
	instructionPtr int
	loopStack      *linkedliststack.Stack
}

// NewInstructionState creates a new tracking state for the given program.
func NewInstructionState(program []byte) *InstructionState {
	return &InstructionState{instructions: program, instructionPtr: 0, loopStack: linkedliststack.New()}
}

// GetInstruction returns the instruction at the current index in the program.
func (s *InstructionState) GetInstruction() byte {
	if s.instructionPtr >= len(s.instructions) {
		return 0
	}
	return s.instructions[s.instructionPtr]
}

// IncrementInstructionPtr moves the program counter for the program up by one.
func (s *InstructionState) IncrementInstructionPtr() {
	s.instructionPtr++
}

// StoreLoopPointer places the index a loop starts at into a stack for retrieval
// later. Used when starting a loop.
func (s *InstructionState) StoreLoopPointer() {
	s.loopStack.Push(s.instructionPtr)
}

// PopLoopPointer pops the most recently store loop index off of the stack. Does not
// set the current program index to its value. Used when terminating a loop.
func (s *InstructionState) PopLoopPointer() {
	s.loopStack.Pop()
}

// RestoreLoopPointer looks at the most recently stored loop index
// and sets the current instruction index to that value. Does not
// pop the index off of the stack.
func (s *InstructionState) RestoreLoopPointer() error {
	val, ok := s.loopStack.Peek()
	if !ok {
		return errors.New("no loop index to jump back to")
	}
	ptrVal, ok := val.(int)
	if !ok {
		return fmt.Errorf("value in stack was not an int, got %v", val)
	}
	s.instructionPtr = ptrVal
	return nil
}

// SkipLoop advances the program past a loop. Ignores nested loops within
// the skipped loop.
func (s *InstructionState) SkipLoop() {
	loopTerms := 1
	for loopTerms > 0 {
		s.instructionPtr++
		if string(s.instructions[s.instructionPtr]) == "[" {
			loopTerms++
		} else if string(s.instructions[s.instructionPtr]) == "]" {
			loopTerms--
		}
	}
}

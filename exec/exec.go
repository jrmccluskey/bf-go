// Package exec contains the logic for running a bf program.
package exec

import (
	"fmt"

	"github.com/jrmccluskey/bf-go/program"
	"github.com/jrmccluskey/bf-go/tape"
)

// Runner encapsulates the program and the data state required to run a bf program.
type Runner struct {
	prog *program.InstructionState
	data *tape.DataState
}

// NewRunner instantiates a new Runner struct.
func NewRunner(prog *program.InstructionState, data *tape.DataState) *Runner {
	return &Runner{prog: prog, data: data}
}

// Execute runs the provided bf program.
func (r *Runner) Execute(input []byte) error {
	inst := r.prog.GetInstruction()
	inputCounter := 0
	inputSize := len(input)
	for inst != 0 {
		switch string(inst) {
		case ">":
			r.data.IncrementPointer()
		case "<":
			r.data.DecrementPointer()
		case "+":
			r.data.IncrementByte()
		case "-":
			r.data.DecrementByte()
		case ".":
			fmt.Print(string(r.data.OutputByte()))
		case ",":
			if inputCounter >= inputSize {
				continue
			}
			r.data.AcceptByte(input[inputCounter])
			inputCounter++
		case "[":
			if r.data.OutputByte() == 0 {
				r.prog.SkipLoop()
			} else {
				r.prog.StoreLoopPointer()
			}
		case "]":
			if r.data.OutputByte() == 0 {
				r.prog.PopLoopPointer()
			} else {
				err := r.prog.RestoreLoopPointer()
				if err != nil {
					return err
				}
			}
		default:
			// Ignore non-instruction characters
		}
		r.prog.IncrementInstructionPtr()
		inst = r.prog.GetInstruction()
	}
	return nil
}

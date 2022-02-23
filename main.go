// Package main arranges and runs a bf file.
package main

import (
	"errors"
	"flag"
	"log"
	"os"
	"strings"

	"github.com/jrmccluskey/bf-go/exec"
	"github.com/jrmccluskey/bf-go/program"
	"github.com/jrmccluskey/bf-go/tape"
)

var (
	programFile = flag.String("program_file", "", "Path to a .bf file to execute (required).")
)

func checkFlags() error {
	if *programFile == "" {
		return errors.New("must provide a .bf file to execute")
	}
	if !strings.HasSuffix(*programFile, ".bf") {
		return errors.New("file must have a .bf extention")
	}
	return nil
}

func main() {
	flag.Parse()
	if err := checkFlags(); err != nil {
		log.Fatalf("failed to execute program, got %v", err)
		return
	}

	prog, err := os.ReadFile(*programFile)
	if err != nil {
		log.Fatalf("failed to read file, got %v", err)
	}

	instState := program.NewInstructionState(prog)
	tapeState := tape.NewDataState()

	progRunner := exec.NewRunner(instState, tapeState)
	if err := progRunner.Execute(); err != nil {
		panic(err)
	}
}

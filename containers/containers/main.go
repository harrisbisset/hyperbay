package main

import (
	"errors"
	"os"
)

func main() {
	if err := handle_input(os.Args); err != nil {
		panic(err) // todo
	}
}

func handle_input(args []string) error {
	var phase Phase
	var err error

	if phase, err = CreatePhase(args); err != nil {
		return err
	} else if phase == nil { // shouldn't happen
		return errors.New("phase unexpectedly nil")
	}

	return phase.RunProcess()
}

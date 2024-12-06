package main

import (
	"flag"
	"fmt"

	"github.com/maximekuhn/brainfuck/cmd/interpreter/interactive"
)

// TODO: handle filepath opt

func main() {
	interactive := flag.Bool("interactive", false, "start interactive interpreter")
	filePath := flag.String("file", "", "interpret the specified Brainfuck file")
	flag.Parse()

	if *interactive {
		if err := runInteractiveInterpretor(); err != nil {
			panic(err)
		}
		return
	}

	if *filePath != "" {
		return
	}

	flag.Usage()
}

func runInteractiveInterpretor() error {
	fmt.Println("starting interactive interpreter...")
	itir := interactive.NewInteractiveInterpreter()
	return itir.Run()
}

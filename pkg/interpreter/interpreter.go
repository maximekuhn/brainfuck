package interpreter

import (
	"errors"
	"fmt"
	"io"

	"github.com/maximekuhn/brainfuck/pkg/parser"
)

const memorySize int = 3_000

type Interpreter struct {
	in     io.Reader
	out    io.Writer
	memory [memorySize]int
	ptr    int
}

func NewInterpreter(in io.Reader, out io.Writer) *Interpreter {
	return &Interpreter{
		in:     in,
		out:    out,
		memory: [memorySize]int{},
		ptr:    0,
	}
}

func (i *Interpreter) Run(ast *parser.Ast) error {
	for _, node := range ast.Statements {
		if err := i.evalNode(node); err != nil {
			return err
		}
	}
	return nil
}

func (i *Interpreter) Dump() ([]int, int) {
	// make a copy to prevent outside modification of interpreter's memory
	tmp := make([]int, memorySize)
	copy(tmp, i.memory[:])
	return tmp, i.ptr
}

func (i *Interpreter) evalNode(node *parser.Node) error {
	switch node.Type {
	case parser.NodeIncrement:
		return i.evalNodeIncrement()
	case parser.NodeDecrement:
		return i.evalNodeDecrement()
	case parser.NodeNext:
		return i.evalNodeNext()
	case parser.NodePrevious:
		return i.evalNodePrevious()
	case parser.NodeOutput:
		return i.evalNodeOutput()
	case parser.NodeInput:
		return i.evalNodeInput()
	case parser.NodeLoop:
		return i.evalNodeLoop(node)
	}
	// unreachable
	return errors.New("unreachable")
}

func (i *Interpreter) evalNodeIncrement() error {
	i.memory[i.ptr]++
	return nil
}

func (i *Interpreter) evalNodeDecrement() error {
	i.memory[i.ptr]--
	return nil
}

func (i *Interpreter) evalNodeNext() error {
	if i.ptr == memorySize-1 {
		return errors.New("ptr is already at max value")
	}
	i.ptr++
	return nil
}

func (i *Interpreter) evalNodePrevious() error {
	if i.ptr == 0 {
		return errors.New("ptr is already at min value")
	}
	i.ptr--
	return nil
}

func (i *Interpreter) evalNodeInput() error {
	b := make([]byte, 1)
	n, err := i.in.Read(b)
	if err != nil {
		if errors.Is(err, io.EOF) {
			// nothing to read
			return nil
		}
		return err
	}
	if n != 1 {
		return fmt.Errorf("in.Read(): expected to read 1 byte but read %d", n)
	}
	val := int(b[0])
	i.memory[i.ptr] = val
	return nil
}

func (i *Interpreter) evalNodeOutput() error {
	val := rune(i.memory[i.ptr])
	_, err := i.out.Write([]byte(string(val)))
	return err
}

func (i *Interpreter) evalNodeLoop(nodeLoop *parser.Node) error {
	for i.memory[i.ptr] != 0 {
		for _, node := range nodeLoop.Child {
			err := i.evalNode(node)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

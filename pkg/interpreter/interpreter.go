package interpreter

import (
	"errors"
	"io"

	"github.com/maximekuhn/brainfuck/pkg/parser"
)

const memorySize int = 3_000

type Interpreter struct {
	ast    *parser.Ast
	in     io.Reader
	out    io.Writer
	memory [memorySize]int
	ptr    int
}

func NewInterpreter(ast *parser.Ast, in io.Reader, out io.Writer) *Interpreter {
	return &Interpreter{
		ast:    ast,
		in:     in,
		out:    out,
		memory: [memorySize]int{},
		ptr:    0,
	}
}

func (i *Interpreter) Run() error {
	for _, node := range i.ast.Statements {
		if err := i.evalNode(node); err != nil {
			return err
		}
	}
	return nil
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
	return errors.New("TODO: implement evalNodeInput")
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
package webapp

import (
	"bytes"
	"io"

	"github.com/maximekuhn/brainfuck/pkg/interpreter"
	"github.com/maximekuhn/brainfuck/pkg/lexer"
	"github.com/maximekuhn/brainfuck/pkg/parser"
)

type service struct {
}

func newService() *service {
	return &service{}
}

func (s *service) runCode(code, inputArgs string) (string, error) {
	lexer := lexer.NewLexer(code)
	toks, err := lexer.Lex()
	if err != nil {
		return "", err
	}

	parser := parser.NewParser(toks)
	ast, err := parser.Parse()
	if err != nil {
		return "", err
	}

	in := newInputReader(inputArgs)
	var out bytes.Buffer
	itp := interpreter.NewInterpreter(in, &out)
	runErr := itp.Run(ast)
	return out.String(), runErr
}

type inputReader struct {
	input []byte
	idx   int
}

func newInputReader(input string) *inputReader {
	return &inputReader{
		input: []byte(input),
		idx:   0,
	}
}

func (ir *inputReader) Read(b []byte) (n int, err error) {
	if ir.idx >= len(ir.input) {
		return 0, io.EOF
	}
	b[0] = ir.input[ir.idx]
	ir.idx++
	return 1, nil
}

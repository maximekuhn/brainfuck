package interpreter

import (
	"bytes"
	_ "embed"
	"io"
	"strings"
	"testing"

	"github.com/maximekuhn/brainfuck/pkg/lexer"
	"github.com/maximekuhn/brainfuck/pkg/parser"
)

//go:embed testdata/helloworld.bf
var srcHelloWorldBf string

func TestInterpreterRun(t *testing.T) {
	testcases := []struct {
		title          string
		input          string
		expectedOutput string
		expectedError  error
	}{
		{
			title:          "Hello world",
			input:          srcHelloWorldBf,
			expectedOutput: "Hello, World!",
			expectedError:  nil,
		},
	}

	for _, test := range testcases {
		t.Run(test.title, func(t *testing.T) {
			ast, err := getAst(test.input)
			if err != nil {
				t.Fatalf("getAst(): expected ok got err %v", err)
			}
			out := bytes.NewBuffer([]byte{})
			i := NewInterpreter(ast, strings.NewReader(test.input), out)
			err = i.Run()
			if test.expectedError != nil && test.expectedError.Error() != err.Error() {
				t.Fatalf("Run(): expected err %v got %v", test.expectedError, err)
			}
			if test.expectedError != nil && test.expectedError.Error() == err.Error() {
				return
			}
			if test.expectedError == nil && err != nil {
				t.Fatalf("Run(): expected ok got err %v", err)
			}
			actualOutput, err := io.ReadAll(out)
			if err != nil {
				t.Fatalf("buffer.ReadAll(): expected ok got err %v", err)
			}
			actualOutputStr := string(actualOutput)
			if test.expectedOutput != actualOutputStr {
				t.Fatalf("Run(): expected output '%s' got '%s'", test.expectedOutput, actualOutputStr)
			}
		})
	}
}

func getAst(input string) (*parser.Ast, error) {
	lexer := lexer.NewLexer(input)
	tokens, err := lexer.Lex()
	if err != nil {
		return nil, err
	}
	parser := parser.NewParser(tokens)
	ast, err := parser.Parse()
	if err != nil {
		return nil, err
	}
	return ast, nil
}

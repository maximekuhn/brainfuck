package interpreter

import (
	"bytes"
	_ "embed"
	"errors"
	"io"
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/maximekuhn/brainfuck/pkg/lexer"
	"github.com/maximekuhn/brainfuck/pkg/parser"
)

//go:embed testdata/helloworld.bf
var srcHelloWorldBf string

//go:embed testdata/helloworldInput.bf
var srcHelloWorldInputBf string

func TestInterpreterRun(t *testing.T) {
	testcases := []struct {
		title          string
		ast            *parser.Ast
		input          string // act as stdin (ASCII)
		expectedOutput string // act as stdout
		expectedError  error
	}{
		{
			title:          "Hello world",
			ast:            getAst(srcHelloWorldBf),
			input:          "",
			expectedOutput: "Hello, World!",
			expectedError:  nil,
		},
		{
			title: "Hello world with input",
			ast:   getAst(srcHelloWorldInputBf),
			// 'Hello, World!' in ASCII
			input:          "72-101-108-108-111-44-32-87-111-114-108-100-33",
			expectedOutput: "Hello, World!",
			expectedError:  nil,
		},
	}

	for _, test := range testcases {
		t.Run(test.title, func(t *testing.T) {
			out := bytes.NewBuffer([]byte{})
			in := newTestStdin(test.input)
			i := NewInterpreter(in, out)
			err := i.Run(test.ast)
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

func getAst(src string) *parser.Ast {
	lexer := lexer.NewLexer(src)
	tokens, err := lexer.Lex()
	if err != nil {
		panic(err)
	}
	parser := parser.NewParser(tokens)
	ast, err := parser.Parse()
	if err != nil {
		panic(err)
	}
	return ast
}

// testStdin is a structure that "mocks" STDIN.
// It contains a list of ASCII values that will be used as input whenever the brainfuck
// program tries to read something from the input source.
type testStdin struct {
	values []int
}

func newTestStdin(input string) *testStdin {
	fragments := strings.Split(input, "-")
	if len(fragments) == 1 && fragments[0] == input {
		return &testStdin{
			values: []int{},
		}
	}
	values := make([]int, 0)
	for _, f := range fragments {
		val, err := strconv.Atoi(f)
		if err != nil {
			panic(err)
		}
		values = append(values, val)
	}
	slices.Reverse(values)
	return &testStdin{
		values: values,
	}
}

func (ts *testStdin) hasNext() bool {
	return len(ts.values) > 0
}

func (ts *testStdin) Read(b []byte) (n int, err error) {
	if !ts.hasNext() {
		return 0, io.EOF
	}
	if len(b) < 1 {
		return 0, errors.New("target buffer is not big enough (expected minimum 1 byte)")
	}
	idx := len(ts.values) - 1
	val := ts.values[idx]
	ts.values = ts.values[:idx]
	b[0] = byte(val)
	return 1, nil

}

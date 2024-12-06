package parser

import (
	"reflect"
	"strings"
	"testing"

	"github.com/maximekuhn/brainfuck/pkg/lexer"
)

func TestParser(t *testing.T) {
	testcases := []struct {
		title         string
		tokens        []lexer.Token
		expectedAst   *Ast
		expectedError error
	}{
		{
			title:         "No input token",
			tokens:        []lexer.Token{},
			expectedAst:   &Ast{},
			expectedError: nil,
		},
		{
			title:  "A single token",
			tokens: []lexer.Token{lexer.TokenIncrement},
			expectedAst: &Ast{Statements: []*Node{
				{
					Type:  NodeIncrement,
					Child: nil,
				}}},
			expectedError: nil,
		},
		{
			title:  "Multiple tokens",
			tokens: []lexer.Token{lexer.TokenIncrement, lexer.TokenDecrement, lexer.TokenOutput},
			expectedAst: &Ast{
				Statements: []*Node{
					{Type: NodeIncrement, Child: nil},
					{Type: NodeDecrement, Child: nil},
					{Type: NodeOutput, Child: nil},
				},
			},
			expectedError: nil,
		},
		{
			title:  "Loop",
			tokens: []lexer.Token{lexer.TokenLoopStart, lexer.TokenIncrement, lexer.TokenLoopEnd},
			expectedAst: &Ast{Statements: []*Node{
				{
					Type: NodeLoop,
					Child: []*Node{
						{Type: NodeIncrement, Child: nil},
					},
				},
			}},
			expectedError: nil,
		},
		{
			title: "Nested loops",
			tokens: []lexer.Token{
				lexer.TokenLoopStart, // -------+
				lexer.TokenIncrement, //        |
				lexer.TokenLoopStart, // ---+   |
				lexer.TokenDecrement, //    |   |
				lexer.TokenLoopEnd,   // ---+   |
				lexer.TokenOutput,    //        |
				lexer.TokenLoopEnd,   // -------+
			},
			expectedAst: &Ast{Statements: []*Node{
				{
					Type: NodeLoop,
					Child: []*Node{
						{Type: NodeIncrement, Child: nil},
						{
							Type: NodeLoop,
							Child: []*Node{
								{Type: NodeDecrement, Child: nil},
							},
						},
						{Type: NodeOutput, Child: nil},
					},
				},
			}},
			expectedError: nil,
		},
		{
			title: "Hello world input program",
			tokens: []lexer.Token{ // ,[>,]<[<]>[.>]
				lexer.TokenInput,                                                            // ,
				lexer.TokenLoopStart, lexer.TokenNext, lexer.TokenInput, lexer.TokenLoopEnd, // [>,]
				lexer.TokenPrevious,                                           // <
				lexer.TokenLoopStart, lexer.TokenPrevious, lexer.TokenLoopEnd, // [<]
				lexer.TokenNext,                                                              // >
				lexer.TokenLoopStart, lexer.TokenOutput, lexer.TokenNext, lexer.TokenLoopEnd, // [.>]

			},
			expectedAst: &Ast{
				Statements: []*Node{
					{Type: NodeInput, Child: nil},
					{Type: NodeLoop, Child: []*Node{
						{Type: NodeNext, Child: nil},
						{Type: NodeInput, Child: nil},
					}},
					{Type: NodePrevious, Child: nil},
					{Type: NodeLoop, Child: []*Node{
						{Type: NodePrevious, Child: nil},
					}},
					{Type: NodeNext, Child: nil},
					{Type: NodeLoop, Child: []*Node{
						{Type: NodeOutput, Child: nil},
						{Type: NodeNext, Child: nil},
					}},
				},
			},
			expectedError: nil,
		},
	}

	for _, test := range testcases {
		t.Run(test.title, func(t *testing.T) {
			p := NewParser(test.tokens)
			actualAst, err := p.Parse()
			if test.expectedError != nil && test.expectedError.Error() != err.Error() {
				t.Fatalf("Parse(): expected err %v got %v", test.expectedError, err)
			}
			if test.expectedError != nil && test.expectedError.Error() == err.Error() {
				return
			}
			if test.expectedError == nil && err != nil {
				t.Fatalf("Parse(): expected ok got err %v", err)
			}
			if len(test.expectedAst.Statements) == 0 && len(actualAst.Statements) == 0 {
				return
			}
			if !reflect.DeepEqual(test.expectedAst, actualAst) {
				t.Errorf("Parse(): expected AST %s got %s", formatAst(test.expectedAst), formatAst(actualAst))
			}
		})
	}

}

func formatAst(ast *Ast) string {
	if ast == nil {
		return "nil"
	}
	var sb strings.Builder
	sb.WriteString("{")
	for _, node := range ast.Statements {
		sb.WriteString(formatNode(node))
		sb.WriteRune(' ')
	}
	sb.WriteString("}")
	return sb.String()
}

func formatNode(n *Node) string {
	if n == nil {
		return "nil"
	}
	var sb strings.Builder
	sb.WriteRune(' ')
	sb.WriteString(formatNodeType(n.Type))
	if len(n.Child) > 0 {
		sb.WriteRune('(')
	}
	for _, c := range n.Child {
		sb.WriteString(formatNode(c))
	}
	if len(n.Child) > 0 {
		sb.WriteRune(')')
	}
	sb.WriteRune(' ')
	return sb.String()
}

func formatNodeType(t NodeType) string {
	switch t {
	case NodeIncrement:
		return "NodeIncrement"
	case NodeDecrement:
		return "NodeDecrement"
	case NodeNext:
		return "NodeNext"
	case NodePrevious:
		return "NodePrevious"
	case NodeOutput:
		return "NodeOutput"
	case NodeInput:
		return "NodeInput"
	case NodeLoop:
		return "NodeLoop"
	}
	return "UnknownNodeType"
}

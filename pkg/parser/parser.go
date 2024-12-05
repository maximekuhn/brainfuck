package parser

import (
	"errors"

	"github.com/maximekuhn/brainfuck/pkg/lexer"
)

// TODO: check for EOF

type Parser struct {
	tokens    []lexer.Token
	currIdx   int
	isInLoop  bool
	loopDepth int
	ns        *nodeStack
}

// NewParser creates a new parser with the tokens as input stream.
func NewParser(tokens []lexer.Token) *Parser {
	return &Parser{
		tokens:    tokens,
		currIdx:   0,
		isInLoop:  false,
		loopDepth: 0,
		ns:        newNodeStack(),
	}
}

// Parse attempts to create an abstract syntax tree (AST)
// from the input stream of tokens.
//
// If an error occured, the returned AST should not be used.
func (p *Parser) Parse() (*Ast, error) {
	ast := &Ast{
		Statements: make([]*Node, 0),
	}
	for p.hasNext() {
		node, err := p.parse()
		if err != nil {
			return ast, err
		}
		if node == nil {
			continue
		}
		loopNode, found := p.ns.peek()
		if !found {
			ast.Statements = append(ast.Statements, node)
			continue
		}
		loopNode.Child = append(loopNode.Child, node)
	}
	return ast, nil
}

func (p *Parser) parse() (*Node, error) {
	next := p.getNext()
	switch next {
	case lexer.TokenIncrement:
		return &Node{Type: NodeIncrement, Child: nil}, nil
	case lexer.TokenDecrement:
		return &Node{Type: NodeDecrement, Child: nil}, nil
	case lexer.TokenNext:
		return &Node{Type: NodeNext, Child: nil}, nil
	case lexer.TokenPrevious:
		return &Node{Type: NodePrevious, Child: nil}, nil
	case lexer.TokenOutput:
		return &Node{Type: NodeOutput, Child: nil}, nil
	case lexer.TokenInput:
		return &Node{Type: NodeInput, Child: nil}, nil
	case lexer.TokenLoopStart:
		loopNode := &Node{
			Type:  NodeLoop,
			Child: make([]*Node, 0),
		}
		p.ns.push(loopNode)
		return nil, nil
	case lexer.TokenLoopEnd:
		loopNode, found := p.ns.pop()
		if !found {
			return nil, errors.New("got TokenLoopEnd but currently not in a loop")
		}
		return loopNode, nil
	}

	// unreachable
	return nil, errors.New("unreachable")
}

func (p *Parser) hasNext() bool {
	return p.currIdx < len(p.tokens)
}

func (p *Parser) getNext() lexer.Token {
	next := p.tokens[p.currIdx]
	p.currIdx++
	return next
}

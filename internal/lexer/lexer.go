package lexer

import "fmt"

type Lexer struct {
	input   string
	currIdx int
}

func NewLexer(input string) *Lexer {
	return &Lexer{
		input:   input,
		currIdx: 0,
	}
}

func (l *Lexer) Lex() ([]Token, error) {
	toks := make([]Token, 0)
	for l.hasNext() {
		next := l.getNext()
		tok, err := toToken(next)
		if err != nil {
			return toks, err
		}
		toks = append(toks, tok)
	}
	return toks, nil
}

func (l *Lexer) hasNext() bool {
	return l.currIdx < len(l.input)
}

func (l *Lexer) getNext() rune {
	next := rune(l.input[l.currIdx])
	l.currIdx++
	return next
}

func toToken(c rune) (Token, error) {
	switch c {
	case '+':
		return TokenIncrement, nil
	case '-':
		return TokenDecrement, nil
	case '>':
		return TokenNext, nil
	case '<':
		return TokenPrevious, nil
	case '.':
		return TokenOutput, nil
	case ',':
		return TokenInput, nil
	case '[':
		return TokenLoopStart, nil
	case ']':
		return TokenLoopEnd, nil
	}
	return TokenNext, fmt.Errorf("unknown token: %c", c)
}

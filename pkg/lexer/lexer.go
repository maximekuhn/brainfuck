package lexer

import (
	"fmt"
	"strings"
)

type Lexer struct {
	input   string
	currIdx int
}

func NewLexer(input string) *Lexer {
	return &Lexer{
		input:   removeGarbageInput(input),
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
	return TokenNext, fmt.Errorf("unknown token: '%c'", c)
}

func removeGarbageInput(input string) string {
	// the current method isn't very optimized as we will read the whole input once
	// and remove all characters that are considered as comments (all except +-<>.,[])
	// maybe one day it will be re-worked but for a toy project like this it is perfectly fine
	var cleaned strings.Builder
	for _, l := range input {
		if isAllowedLitteral(l) {
			cleaned.WriteRune(l)
		}
	}
	return cleaned.String()
}

func isAllowedLitteral(l rune) bool {
	var allowedLitterals = []rune{'+', '-', '<', '>', '.', ',', '[', ']'}
	for _, al := range allowedLitterals {
		if l == al {
			return true
		}
	}
	return false
}

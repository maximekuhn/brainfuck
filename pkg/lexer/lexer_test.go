package lexer

import (
	"reflect"
	"strings"
	"testing"
)

func TestLexer(t *testing.T) {
	testcases := []struct {
		title          string
		input          string
		expectedTokens []Token
		expectedError  error
	}{
		{
			title:          "Empty input",
			input:          "",
			expectedTokens: []Token{},
			expectedError:  nil,
		},
		{
			title: "All tokens",
			input: "><+-.,[]",
			expectedTokens: []Token{
				TokenNext, TokenPrevious,
				TokenIncrement, TokenDecrement,
				TokenOutput, TokenInput,
				TokenLoopStart, TokenLoopEnd},
			expectedError: nil,
		},
		{
			title: "Invalid token should be considered as a comment",
			input: "<>+/[",
			expectedTokens: []Token{
				TokenPrevious, TokenNext,
				TokenIncrement, TokenLoopStart,
			},
			expectedError: nil,
		},
		{
			title:          "Only one token",
			input:          ".",
			expectedTokens: []Token{TokenOutput},
			expectedError:  nil,
		},
	}

	for _, test := range testcases {
		t.Run(test.title, func(t *testing.T) {
			l := NewLexer(test.input)
			actualTokens, err := l.Lex()
			if test.expectedError != nil && err != nil && test.expectedError.Error() != err.Error() {
				t.Fatalf("Lex(): expected err %v got %v", test.expectedError, err)
			}
			if test.expectedError != nil && err != nil && test.expectedError.Error() == err.Error() {
				return
			}
			if test.expectedError != nil && err == nil {
				t.Fatalf("Lex(): expected err %v got ok", test.expectedError)
			}
			if !reflect.DeepEqual(test.expectedTokens, actualTokens) {
				t.Errorf("Lex(): expected tokens %s got %s", formatTokens(test.expectedTokens), formatTokens(actualTokens))
			}
		})

	}
}

func formatTokens(tokens []Token) string {
	var sb strings.Builder
	sb.WriteRune('[')
	for _, token := range tokens {
		switch token {
		case TokenIncrement:
			sb.WriteString("TokenIncrement")
		case TokenDecrement:
			sb.WriteString("TokenDecrement")
		case TokenNext:
			sb.WriteString("TokenNext")
		case TokenPrevious:
			sb.WriteString("TokenPrevious")
		case TokenOutput:
			sb.WriteString("TokenOutput")
		case TokenInput:
			sb.WriteString("TokenInput")
		case TokenLoopStart:
			sb.WriteString("TokenLoopStart")
		case TokenLoopEnd:
			sb.WriteString("TokenLoopEnd")
		default:
			sb.WriteString("UnknownToken")
		}
		sb.WriteRune(' ')
	}
	sb.WriteRune(']')
	return sb.String()
}

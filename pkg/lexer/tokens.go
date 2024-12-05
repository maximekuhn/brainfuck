package lexer

// TODO: TokenEOF (?)

type Token string

const (
	TokenIncrement Token = "+"
	TokenDecrement Token = "-"
	TokenNext      Token = ">"
	TokenPrevious  Token = "<"
	TokenOutput    Token = "."
	TokenInput     Token = ","
	TokenLoopStart Token = "["
	TokenLoopEnd   Token = "]"
)

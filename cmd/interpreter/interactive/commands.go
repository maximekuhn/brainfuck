package interactive

type commandType int

const (
	commandDump  commandType = iota // dump the interpreter state
	commandReset                    // reset the interpreter
	commandExit                     // exit interpreter
	commandRun                      // run the following brainfuck code
	commandHelp                     // print help
)

type command struct {
	cmdType commandType
	opts    interface{}
}

type commandDumpOpts struct {
	memorySize int // number of bytes to dump from interpreter's memory
	offset     int
}

type commandRunOpts struct {
	brainfuckCode string
}

func (c command) String() string {
	switch c.cmdType {
	case commandDump:
		return "commandDump"
	case commandReset:
		return "commandReset"
	case commandExit:
		return "commandExit"
	}
	return "unknownCommandType"
}

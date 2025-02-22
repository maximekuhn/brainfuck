package interactive

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/maximekuhn/brainfuck/pkg/interpreter"
	"github.com/maximekuhn/brainfuck/pkg/lexer"
	"github.com/maximekuhn/brainfuck/pkg/parser"
)

type InteractiveInterpreter struct {
	itp *interpreter.Interpreter
	ir  *inputReader
}

func NewInteractiveInterpreter() *InteractiveInterpreter {
	inputReader := inputReader{}
	itp := interpreter.NewInterpreter(&inputReader, os.Stdout)
	return &InteractiveInterpreter{
		itp: itp,
		ir:  &inputReader,
	}
}

func (i *InteractiveInterpreter) Run() error {
	fmt.Println("type `help` to see the list of available commands")
	scanner := bufio.NewScanner(os.Stdin)
	i.ir.setScanner(scanner)
	prompt()
	for scanner.Scan() {
		input := scanner.Text()
		cmd, err := parseCommand(input)
		if err != nil {
			fmt.Printf("error parsing input: %s\n", err)
			prompt()
			continue
		}
		if cmd.cmdType == commandExit {
			fmt.Println("bye bye...")
			return nil
		}
		if err := i.handleCmd(cmd); err != nil {
			fmt.Printf("error handling command: %s\n", err)
			prompt()
			continue
		}
		fmt.Println()
		prompt()
	}
	return nil
}

func (i *InteractiveInterpreter) handleCmd(cmd *command) error {
	switch cmd.cmdType {
	case commandDump:
		return i.handleCmdDump(cmd)
	case commandRun:
		return i.handleCmdRun(cmd)
	case commandHelp:
		i.handleCmdHelp()
		return nil
	default:
		// unreachable
		return fmt.Errorf("command '%s' can not be handled", cmd)
	}
}

func (i *InteractiveInterpreter) handleCmdDump(cmd *command) error {
	if cmd.cmdType != commandDump {
		return errors.New("wtf")
	}
	opts, ok := cmd.opts.(commandDumpOpts)
	if !ok {
		return errors.New("wtf")
	}
	mem, ptr := i.itp.Dump()
	startIdx := opts.offset
	endIdx := opts.offset + opts.memorySize
	fmt.Printf("mem: %v\n", mem[startIdx:endIdx])
	fmt.Printf("ptr: %v\n", ptr)
	return nil
}

func (i *InteractiveInterpreter) handleCmdRun(cmd *command) error {
	if cmd.cmdType != commandRun {
		return errors.New("wtf")
	}
	opts, ok := cmd.opts.(commandRunOpts)
	if !ok {
		return errors.New("wtf")
	}

	toks, err := lexer.NewLexer(opts.brainfuckCode).Lex()
	if err != nil {
		return err
	}
	ast, err := parser.NewParser(toks).Parse()
	if err != nil {
		return err
	}
	return i.itp.Run(ast)
}

func (i *InteractiveInterpreter) handleCmdHelp() {
	fmt.Println(`
        * run <Brainfuck code>                             - run the provided code
        * dump [size] [offset]                             - dump interpreter's state
        * reset                                            - reset interpreter's state
        * exit                                             - quit the interactive interpreter
        * help                                             - show this menu
        `)
}

func prompt() {
	fmt.Print("~>")
}

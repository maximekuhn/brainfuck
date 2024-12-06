package interactive

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	defaultDumpMemSize int = 10
	defaultDumpOffset  int = 0
)

func parseCommand(input string) (*command, error) {
	if input == "" {
		return nil, errors.New("empty input")
	}
	if input == "exit" || input == "quit" {
		return &command{cmdType: commandExit, opts: nil}, nil
	}
	if input == "reset" {
		return &command{cmdType: commandReset, opts: nil}, nil
	}
	if input == "help" {
		return &command{cmdType: commandHelp, opts: nil}, nil
	}

	fragments := strings.Fields(input)
	cmdName := fragments[0]
	if cmdName == "dump" {
		return parseDumpCommand(fragments)
	}
	if cmdName == "run" {
		return parseRunCommand(fragments)
	}
	return nil, fmt.Errorf("unknown command: '%s'", cmdName)
}

func parseRunCommand(fragments []string) (*command, error) {
	if len(fragments) < 2 {
		return nil, errors.New("run command expect to be followed by some brainfuck code")
	}
	bfCode := strings.Join(fragments[1:], "")
	return &command{cmdType: commandRun, opts: commandRunOpts{brainfuckCode: bfCode}}, nil
}

func parseDumpCommand(fragments []string) (*command, error) {
	opts := commandDumpOpts{
		memorySize: defaultDumpMemSize,
		offset:     defaultDumpOffset,
	}
	if len(fragments) < 2 {
		return &command{cmdType: commandDump, opts: opts}, nil
	}
	if len(fragments) >= 2 {
		memSizeStr := fragments[1]
		memSize, err := strconv.Atoi(memSizeStr)
		if err != nil {
			return nil, err
		}
		opts.memorySize = memSize
	}
	if len(fragments) >= 3 {
		offsetStr := fragments[2]
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			return nil, err
		}
		opts.offset = offset
	}
	return &command{cmdType: commandDump, opts: opts}, nil
}

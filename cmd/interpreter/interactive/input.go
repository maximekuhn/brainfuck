package interactive

import (
	"bufio"
	"errors"
	"fmt"
	"io"
)

// inputReader implements io.Read and prompts user for an input.
// It always read at most 1 byte.
type inputReader struct {
	scanner *bufio.Scanner
}

func (ir *inputReader) setScanner(s *bufio.Scanner) {
	ir.scanner = s
}

func (ir *inputReader) Read(b []byte) (n int, err error) {
	if ir.scanner == nil {
		return 0, errors.New("scanner has not be initalized")
	}

	fmt.Print("input: ")
	ir.scanner.Scan()
	input := ir.scanner.Bytes()
	if len(input) == 0 {
		return 0, io.EOF
	}
	b[0] = input[0]
	return 1, nil
}

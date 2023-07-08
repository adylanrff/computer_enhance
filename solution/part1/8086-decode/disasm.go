package i8086

import (
	"errors"
	"fmt"
	"io"
	"os"
)

type Disassembler struct {
	Filepath string
}

var (
	ErrorEmptyFilepath = errors.New("empty filepath")
	ErrorInvalidFile   = errors.New("invalid file")
)

func (d *Disassembler) Disassemble() error {
	if d.Filepath == "" {
		return ErrorEmptyFilepath
	}

	f, err := os.Open(d.Filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	buffer := make([]byte, 2)

	// 16 bits for asm
	fmt.Println("bits 16")

	for {
		n, err := f.Read(buffer)
		if n == 0 {
			return nil
		}

		if n < 2 {
			return ErrorInvalidFile
		}

		if err != nil {
			if errors.Is(err, io.EOF) {
				// Assume done loading
				return nil
			}
			return err
		}

		instruction, err := GetInstruction(buffer)
		if err != nil {
			return err
		}

		fmt.Println(instruction.ToAsm())
	}
}

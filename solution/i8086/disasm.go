package i8086

import (
	"encoding/binary"
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

func GetInstruction(b []byte) (Instruction, error) {
	if len(b) != 2 {
		return Instruction{}, errors.New("invalid instruction")
	}

	byteInt := binary.BigEndian.Uint16(b)

	opcodeByte := (byteInt & OpcodeMask) >> OpcodeShift
	direction := (byteInt & DirectionMask) >> DirectionShift
	wide := (byteInt & WideMask) >> WideShift
	memoryMode := (byteInt & MemoryModeMask) >> MemoryModeShift
	reg := (byteInt & RegMask) >> RegShift
	rm := (byteInt & RmMask) >> RmShift

	return Instruction{
		Opcode:    OpcodeMapping[byte(opcodeByte)],
		Direction: direction != 0,
		Wide:      wide != 0,
		Mod:       MemoryMode(memoryMode),
		Reg:       Register(reg),
		RM:        RegisterOrMemory(rm),
	}, nil
}

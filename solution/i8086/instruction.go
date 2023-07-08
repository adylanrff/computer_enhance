package i8086

import (
	"fmt"
)

// Opcodes
type Opcode string

const (
	OpcodeMov Opcode = "mov"
)

var OpcodeMapping = map[byte]Opcode{
	0b100010: OpcodeMov,
}

const (
	OpcodeMask  = 0xFCFF
	OpcodeShift = 10
)

// Direction

type Direction bool

const (
	DirectionFromRegister Direction = false
	DirectionToRegister   Direction = true
)

const (
	DirectionMask  = 0x02FF
	DirectionShift = 9
)

// Wide

const (
	WideMask  = 0x01FF
	WideShift = 8
)

//  MemoryMode

type MemoryMode uint8

const (
	MemoryModeNoDisplacement MemoryMode = iota
	MemoryMode8BitDisplacement
	MemoryMode16BitDisplacement
	MemoryModeRegisterMode
)

const (
	MemoryModeMask  = 0xC0
	MemoryModeShift = 6
)

// Register
type Register uint8

const (
	RegMask  = 0x38
	RegShift = 3
)

var RegisterMapping = map[byte][]string{
	0b000: {"al", "ax"},
	0b001: {"cl", "cx"},
	0b010: {"dl", "dx"},
	0b011: {"bl", "bx"},
	0b100: {"ah", "sp"},
	0b101: {"ch", "bp"},
	0b110: {"dh", "si"},
	0b111: {"bh", "di"},
}

func (r Register) String(wide bool) string {
	var bitset int
	if wide {
		bitset = 1
	}
	return RegisterMapping[byte(r)][bitset]
}

// RM
type RegisterOrMemory uint8

const (
	RmMask  = 0x07
	RmShift = 0
)

func (r RegisterOrMemory) String(wide bool) string {
	var bitset int
	if wide {
		bitset = 1
	}
	return RegisterMapping[byte(r)][bitset]
}

type Instruction struct {
	Opcode    Opcode
	Direction Direction
	Wide      bool
	Mod       MemoryMode
	Reg       Register
	RM        RegisterOrMemory
}

func (i *Instruction) ToAsm() string {
	destStr := i.Reg.String(i.Wide)
	sourceStr := i.RM.String(i.Wide)

	if i.Direction == DirectionFromRegister {
		destStr, sourceStr = sourceStr, destStr
	}

	return fmt.Sprintf("%s %s, %s", i.Opcode, destStr, sourceStr)
}

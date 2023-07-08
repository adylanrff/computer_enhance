package main

import (
	"fmt"
	"os"

	i8086 "github.com/adylanrff/computer_enhance/solution/part1/8086-decode"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("not enough arguments. pass with 8086 binary")
		os.Exit(1)
	}

	filepath := os.Args[1]
	disassembler := i8086.Disassembler{Filepath: filepath}
	if err := disassembler.Disassemble(); err != nil {
		panic(err)
	}
}

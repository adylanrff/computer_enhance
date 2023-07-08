package main

import (
	"fmt"
	"os"

	"github.com/adylanrff/computer_enhance/solution/i8086"
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

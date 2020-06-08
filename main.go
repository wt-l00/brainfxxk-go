package main

import (
	"bufio"
	"log"
	"os"

	"github.com/wt-l00/brainfxxk-go/interpreter"
	"github.com/wt-l00/brainfxxk-go/parser"
)

func main() {
	f, err := os.Open("example/sample")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	var program string

	for {
		scanned := scanner.Scan()
		if !scanned {
			break
		}
		program += scanner.Text()
	}

	p := parser.New(program)
	p.ParseProgram()

	println(interpreter.InterpreterProgram(p.Instructions))
}

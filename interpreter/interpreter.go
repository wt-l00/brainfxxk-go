package interpreter

import (
	"bufio"
	"bytes"
	"os"
)

const (
	memorySize int64 = 300000
)

func InterpreterProgram(instructions []byte) string {
	memory := make([]int, memorySize)
	dataPtr := 0

	var res bytes.Buffer

	for pc := 0; pc < len(instructions); pc++ {
		//println(instructions[pc])
		switch instructions[pc] {
		case '>':
			dataPtr++
		case '<':
			dataPtr--
		case '+':
			memory[dataPtr]++
		case '-':
			memory[dataPtr]--
		case '.':
			res.WriteString(string(memory[dataPtr]))
		case ',':
			scanner := bufio.NewScanner(os.Stdin)
			scanned := scanner.Scan()
			if !scanned {
				return "failed"
			}
			line := scanner.Text()
			char := line[0]
			memory[dataPtr] = int(char)
		case '[':
			bracketNesting := 1
			if memory[dataPtr] == 0 {
				pc++
				for bracketNesting > 0 && pc < len(instructions) {
					if instructions[pc] == ']' {
						bracketNesting--
					} else if instructions[pc] == '[' {
						bracketNesting++
					}
					pc++
				}

				if bracketNesting != 0 {
					return "missed ["
				}
			}
		case ']':
			bracketNesting := 1
			if memory[dataPtr] != 0 {
				for bracketNesting > 0 && pc > 0 {
					pc--
					if instructions[pc] == '[' {
						bracketNesting--
					} else if instructions[pc] == ']' {
						bracketNesting++
					}
				}
				if bracketNesting != 0 {
					return "missed ]"
				}
			}

		default:
			return "failed"
		}
	}

	return res.String()
}

package interpreter

import (
	"bytes"
)

const (
	memorySize int64 = 30000
)

func interpreterProgram(instructions []byte) string {
	memory := make([]int64, memorySize)
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
		//case ',':

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

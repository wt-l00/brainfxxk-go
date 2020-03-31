package interpreter

import (
	"testing"

	"github.com/wt-l00/brainfxxk-go/parser"
)

func TestInterpreterProgram(t *testing.T) {
	interpreterTests := []struct {
		input    string
		expected string
	}{
		{"++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.", "Hello World!\n"},
	}

	for _, program := range interpreterTests {
		p := parser.New(program.input)
		p.ParseProgram()

		if interpreterProgram(p.Instructions) != program.expected {
			t.Errorf("expected=%q, got=%q", program.expected, interpreterProgram(p.Instructions))
		}
	}
}

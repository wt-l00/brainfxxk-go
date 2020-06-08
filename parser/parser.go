package parser

import (
	"bytes"
)

type Program struct {
	code         string
	Instructions []byte
}

func New(input string) *Program {
	p := &Program{code: input}
	return p
}

func (p *Program) ParseProgram() {

	for i := 0; i < len(p.code); i++ {
		switch p.code[i] {
		case '>':
			p.Instructions = append(p.Instructions, '>')
		case '<':
			p.Instructions = append(p.Instructions, '<')
		case '+':
			p.Instructions = append(p.Instructions, '+')
		case '-':
			p.Instructions = append(p.Instructions, '-')
		case '.':
			p.Instructions = append(p.Instructions, '.')
		case ',':
			p.Instructions = append(p.Instructions, ',')
		case '[':
			p.Instructions = append(p.Instructions, '[')
		case ']':
			p.Instructions = append(p.Instructions, ']')
		}
	}
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Instructions {
		out.WriteString(string(s))
	}

	return out.String()
}

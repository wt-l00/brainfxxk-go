package parser

import "bytes"

type program struct {
	instructions []byte
}

func (p *program) parseProgram(code string) {
	for i := 0; i < len(code); i++ {
		switch code[i] {
		case '>':
			p.instructions = append(p.instructions, '>')
		case '<':
			p.instructions = append(p.instructions, '<')
		case '+':
			p.instructions = append(p.instructions, '+')
		case '-':
			p.instructions = append(p.instructions, '-')
		case '.':
			p.instructions = append(p.instructions, '.')
		case ',':
			p.instructions = append(p.instructions, ',')
		case '[':
			p.instructions = append(p.instructions, '[')
		case ']':
			p.instructions = append(p.instructions, ']')
		}
	}
}

func (p *program) String() string {
	var out bytes.Buffer

	for _, s := range p.instructions {
		out.WriteString(string(s))
	}

	return out.String()
}

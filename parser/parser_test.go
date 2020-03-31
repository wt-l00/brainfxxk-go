package parser

import (
	"testing"
)

func TestParseProgram(t *testing.T) {
	parseTests := []struct {
		input    string
		expected string
	}{
		{"><><,,", "><><,,"},
		{">>,xx", ">>,"},
		{">>>------,", ">>>------,"},
	}

	for _, code := range parseTests {
		p := New(code.input)
		p.ParseProgram()
		if p.String() != code.expected {
			t.Errorf("expected=%q, got=%q", code.expected, p.String())
		}
	}
}

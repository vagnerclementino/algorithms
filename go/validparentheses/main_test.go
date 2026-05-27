package main

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
		{"(", false},
		{")", false},
		{"({})", true},
		{"({[()]})", true},
		{"{[}]", false},
		{"", true},
		{"{}[]", true},
		{"[][]{{}}", true},
		{"((()))", true},
		{"{[()()]}", true},
		{"{[(])}", false},
		{"[({})](())", true},
		{"[", false},
		{"]", false},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := isValid(test.input)
			if result != test.expected {
				t.Errorf("isValid(%q) = %v; want %v", test.input, result, test.expected)
			}
		})
	}
}

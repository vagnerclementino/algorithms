package main

import "fmt"

type stack struct {
	data []rune
}

func NewStack() *stack {
	return &stack{
		data: []rune{},
	}
}

func (s *stack) Push(r rune) {
	s.data = append(s.data, r)
}

func (s *stack) Pop() rune {
	r := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return r
}

func (s *stack) Peek() rune {
	if len(s.data) == 0 {
		return '*'
	}
	return s.data[len(s.data)-1]
}

func (s *stack) Size() int {
	return len(s.data)
}

func isOpenPar(r rune) bool {
	switch r {
	case '{':
		return true
	case '(':
		return true
	case '[':
		return true
	default:
		return false
	}
}

func getClosePar(r rune) rune {
	switch r {
	case '{':
		return '}'
	case '(':
		return ')'
	case '[':
		return ']'
	default:
		return '*'
	}
}
func isValid(s string) bool {
	if len(s) == 0 {
		return true
	} else if len(s) == 1 {
		return false
	}
	stack := NewStack()
	for _, r := range s {
		if isOpenPar(r) {
			stack.Push(r)
		} else if r == getClosePar(stack.Peek()) {
			stack.Pop()
		}
	}
	return stack.Size() == 0
}

func main() {
	inputs := []string{
		"()",         // Valid
		"()[]{}",     // Valid
		"(]",         // Invalid
		"([)]",       // Invalid
		"{[]}",       // Valid
		"(",          // Invalid
		")",          // Invalid
		"({})",       // Valid
		"({[()]})",   // Valid
		"{[}]",       // Invalid
		"",           // Valid (Empty string)
		"{}[]",       // Valid
		"[][]{{}}",   // Valid
		"((()))",     // Valid
		"{[()()]}",   // Valid
		"{[(])}",     // Invalid
		"[({})](())", // Valid
		"[",          // Invalid
		"]",          // Invalid
	}
	for _, input := range inputs {
		if isValid(input) {
			fmt.Println("Valid")
		} else {
			fmt.Println("Invalid")
		}

	}
}

package main

import "fmt"

type stack[T any] struct {
	data []T
}

func NewStack[T any]() *stack[T] {
	return &stack[T]{
		data: []T{},
	}
}

func (s *stack[T]) Push(r T) {
	s.data = append(s.data, r)
}

func (s *stack[T]) Pop() T {
	r := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return r
}

func (s *stack[T]) Peek() *T {
	if len(s.data) == 0 {
		return nil
	}
	return &s.data[len(s.data)-1]
}

func (s *stack[T]) Size() int {
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

func getClosePar(r *rune) rune {
	if r == nil {
		return '*'
	}
	switch *r {
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
	stack := NewStack[rune]()
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

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

func (s *stack[T]) Peek() (T, bool) {
	if len(s.data) == 0 {
		var zero T
		return zero, false
	}
	return s.data[len(s.data)-1], true
}

func (s *stack[T]) Size() int {
	return len(s.data)
}

func isOpenPar(r rune) bool {
	switch r {
	case '{', '(', '[':
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

	stack := NewStack[rune]()
	for _, r := range s {
		if isOpenPar(r) {
			stack.Push(r)
		} else {
			top, ok := stack.Peek()
			if !ok || r != getClosePar(top) {
				return false
			}
			stack.Pop()
		}
	}
	return stack.Size() == 0
}

func main() {
	// Test cases
	tests := []string{
		"()",     // true
		"()[]{}", // true
		"(]",     // false
		"([)]",   // false
		"{[]}",   // true
		"(",      // false
		")",      // false
		"([]{})", // true
		"([{}])", // true
		"{[}]",   // false
	}

	for _, test := range tests {
		fmt.Printf("Input: %q, Is Valid: %v\n", test, isValid(test))
	}
}

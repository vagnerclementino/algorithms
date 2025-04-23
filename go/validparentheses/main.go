package main

import "fmt"

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
	stack := []rune{}
	stackSize := 0

	for _, r := range s {
		if isOpenPar(r) {
			stack = append(stack, r)
			stackSize++
		} else if r == getClosePar(stack[stackSize-1]) {
			stack = stack[:stackSize-1]
			stackSize--
		}
	}
	return stackSize == 0
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

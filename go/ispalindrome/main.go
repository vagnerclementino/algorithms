package main

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	cleanedStr := cleanString(s)
	if len(cleanedStr) == 0 { // Handle empty or non-alphanumeric strings
		return true
	}

	start, end := 0, len(cleanedStr)-1
	for start < end {
		if cleanedStr[start] == cleanedStr[end] {
			start++
			end--
		} else {
			return false
		}
	}
	return true
}

func cleanString(s string) []rune {
	var builder strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			builder.WriteRune(unicode.ToLower(r))
		}
	}
	return []rune(builder.String())
}

func main() {
	inputs := []string{
		"A man, a plan, a canal: Panama",
		"race a car",
		"",
		" ",
		"a",
	}

	for _, input := range inputs {
		println(isPalindrome(input))
	}
}

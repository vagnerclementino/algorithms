package main

import (
	"strconv"
)

func decompress(s string) string {
	leftBracketIdx := 0
	rightBracketIdx := 0
	for i := 0; i < len(s); i++ {
		if rune(s[i]) == '[' {
			leftBracketIdx = i
		}
		if rune(s[i]) == ']' {
			rightBracketIdx = i
			return decompressItem(s[:i+1], leftBracketIdx, rightBracketIdx) + decompress(s[i+1:])
		}
	}
	return s

}

func decompressItem(s string, leftBracketIdx, rightBracketIdx int) string {
	if m, err := strconv.Atoi(s[0:leftBracketIdx]); err == nil {
		var decompressStr string
		for i := 0; i < m; i++ {
			decompressStr = decompressStr + s[leftBracketIdx+1:rightBracketIdx]
		}
		return decompressStr
	}
	return s
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecompress(t *testing.T) {
	tests := []struct {
		name     string
		scenario func(t *testing.T)
	}{
		{
			name: "should decompress '3[abc]4[ab]c'",
			scenario: func(t *testing.T) {
				assert.Equal(t, "abcabcabcababababc", decompress("3[abc]4[ab]c"))
			},
		},
		{
			name: "should decompress '5[a]'",
			scenario: func(t *testing.T) {
				assert.Equal(t, "aaaaa", decompress("5[a]"))
			},
		},
				{
			name: "should decompress '2[aabc]d'",
			scenario: func(t *testing.T) {
				assert.Equal(t, "aabcaabcd", decompress("2[aabc]d"))
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.scenario(t)
		})
	}
}

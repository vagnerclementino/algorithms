package palindrome

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestLongestPalindrome(t *testing.T) {
    testCases := []struct {
        input    string
        expected string
    }{
        {"babad", "bab"},
        {"cbbd", "bb"},
        {"a", "a"},
        {"ac", "a"},
        {"bb", "bb"},
    }
    for _, tc := range testCases {
        result := LongestPalindrome(tc.input)
        assert.Equal(t, tc.expected, result, "Unexpected result for input %v", tc.input)
    }
}

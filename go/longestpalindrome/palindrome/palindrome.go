package palindrome

func LongestPalindrome(s string) string {
	longestPalin := string(s[0])
	for i := range s {
		for _, j := range [2]int{i, i + 1} {
			l, r := i, j
			for l >= 0 && r < len(s) && s[l] == s[r] {
				if (r - l + 1) > len(longestPalin) {
					longestPalin = s[l : r+1]
				}
				l--
				r++
			}
		}
	}
	return longestPalin
}

package dividenconc

// finds the length of the longest palindromic
// subsequence in a string
func LPS(s string) int {
	length := lps(s, 0, len(s)-1)
	return length
}

func lps(s string, start int, end int) int {
	if start > end {
		// more than half the string has been traversed
		return 0
	}
	if start == end {
		return 1
	}
	c1 := 0
	if s[start] == s[end] {
		c1 = 2 + lps(s, start+1, end-1)
	}
	c2 := lps(s, start+1, end) // increase 1st string
	c3 := lps(s, start, end-1) // increase 2nd string

	return Max(c1, c2, c3)
}

package dividenconc

// finds the length of the longest palindromic
// continuous substring in a string, s
func LPSubstring(s string) int {
	length := lpsubstring(s, 0, len(s)-1)
	return length
}

func lpsubstring(s string, start int, end int) int {
	if start > end {
		// more than half the string has been traversed
		return 0
	}
	if start == end {
		return 1
	}
	c1 := 0
	if s[start] == s[end] {
		remaining_length := end - start - 1
		if remaining_length == lpsubstring(s, start+1, end-1) {
			c1 = remaining_length + 2
		}
	}
	c2 := lpsubstring(s, start+1, end) // increase 1st string
	c3 := lpsubstring(s, start, end-1) // increase 2nd string

	return Max(c1, c2, c3)
}

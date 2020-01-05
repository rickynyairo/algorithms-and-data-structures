package dividenconc

// finds the length of the longest common
// subsequence in the two strings
func LCS(s1 string, s2 string) int {
	length := lcs(s1, s2, 0, 0)
	return length
}

func lcs(s1 string, s2 string, i1 int, i2 int) int {
	if i1 == len(s1) || i2 == len(s2) {
		// no more characters to compare
		return 0
	}
	c3 := 0
	if string(s1[i1]) == string(s2[i2]) {
		return 1 + lcs(s1, s2, i1+1, i2+1)
	}
	c1 := lcs(s1, s2, i1+1, i2) // increase 1st string
	c2 := lcs(s1, s2, i1, i2+1) // increase 2nd string

	// fmt.Println("We're here ====>", string(s1[i1]), string(s2[i2]))
	return Max(c1, c2, c3)
}

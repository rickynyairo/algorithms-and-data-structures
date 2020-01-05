package dividenconc

// calculates the minimum number of operations required to
// convert string s1 to string s2
func FindMinOperations(s1 string, s2 string) int {
	operations := findMin(s1, s2, 0, 0)
	return operations
}

// returns the minimum of the provided values
func min(values ...int) int {
	minimum := values[0]
	for _, val := range values {
		if val < minimum {
			minimum = val
		}
	}
	return minimum
}

func findMin(s1 string, s2 string, i1 int, i2 int) int {
	if i1 == len(s1) {
		// end of s1, delete all characters in s2
		return len(s2) - i2
	}
	if i2 == len(s2) {
		// end of s2, move all characters in s2 to s2
		return len(s1) - i1
	}
	if string(s1[i1]) == string(s2[i2]) {
		return findMin(s1, s2, i1+1, i2+1)
	}
	c1 := 1 + findMin(s1, s2, i1+1, i2)   // insert
	c2 := 1 + findMin(s1, s2, i1, i2+1)   // delete
	c3 := 1 + findMin(s1, s2, i1+1, i2+1) // replace

	// fmt.Println("We're here ====>", string(s1[i1]), string(s2[i2]))
	return min(c1, c2, c3)
}

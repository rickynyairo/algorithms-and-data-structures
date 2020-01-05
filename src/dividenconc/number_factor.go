package dividenconc

func NumberFactor(n int) int {
	if n == 0 || n == 1 || n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	sub1 := NumberFactor(n - 1)
	sub2 := NumberFactor(n - 2)
	sub3 := NumberFactor(n - 3)
	return sub1 + sub2 + sub3
}

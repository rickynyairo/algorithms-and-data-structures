package dividenconc

func Fibonacci(n int) int {
	if n == 0 {
		panic("Cannot get zeroeth element")
	}
	if n > 0 && n <= 2 {
		return n - 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

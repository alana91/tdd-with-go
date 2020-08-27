package fibonacci

// Fibonacci calculates the number in fibonacci sequence in position n
func Fibonacci(n int) int {
	if n == 0 {
		return 0
	}

	if n <= 2 {
		return 1
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

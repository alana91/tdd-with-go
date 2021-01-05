package fibonacci

// Fibonacci calculates the number in fibonacci sequence in position n
func Fibonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

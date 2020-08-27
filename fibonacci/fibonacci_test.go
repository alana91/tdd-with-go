package fibonacci_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tdd-with-go/fibonacci"
)

func TestFibonacci(t *testing.T) {
	cases := map[int]int{
		0: 0,
		1: 1,
		2: 1,
		3: 2,
	}

	for n, expected := range cases {
		assert.Equal(t, expected, fibonacci.Fibonacci(n))
	}
}

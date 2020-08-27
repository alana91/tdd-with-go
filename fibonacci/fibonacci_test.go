package fibonacci_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tdd-with-go/fibonacci"
)

func TestFibonacci(t *testing.T) {
	assert.Equal(t, fibonacci.Fibonacci(0), 0)
	assert.Equal(t, fibonacci.Fibonacci(1), 1)
	assert.Equal(t, fibonacci.Fibonacci(2), 1)
}

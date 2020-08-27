package fibonacci_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	assert.Equal(t, fibonacci.Fibonacci(0), 0)
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var Tests = []struct {
	name     string
	request  int
	expected int
}{
	{
		name:     "Factorial(0)",
		request:  0,
		expected: 1,
	},
	{
		name:     "Factorial(1)",
		request:  1,
		expected: 1,
	},
	{
		name:     "Factorial(2)",
		request:  2,
		expected: 2,
	},
	{
		name:     "Factorial(3)",
		request:  3,
		expected: 6,
	},
	{
		name:     "Factorial(4)",
		request:  4,
		expected: 24,
	},
	{
		name:     "Factorial(5)",
		request:  5,
		expected: 120,
	},
}

func TestFactorial(t *testing.T) {
	for _, test := range Tests {
		result := Factorial(test.request)
		assert.Equal(t, test.expected, result)
	}
}

func TestRecursiveFactorial(t *testing.T) {
	for _, test := range Tests {
		result := RecursiveFactorial(test.request)
		assert.Equal(t, test.expected, result)
	}
}

func TestTailRecursive(t *testing.T) {
	for _, test := range Tests {
		result := TailRecursive(1, test.request)
		assert.Equal(t, test.expected, result)
	}
}

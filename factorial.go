package main

// A. Factorial dengan Looping
func Factorial(value int) int {
	if value <= 0 {
		return 1
	}

	var result = 1
	for i := value; i >= 1; i-- {
		result = result * i
	}

	return result
}

// B. Factorial dengan recursive
func RecursiveFactorial(value int) int {
	if value <= 0 {
		return 1
	} else {
		return value * RecursiveFactorial(value-1)
	}
}

// C. Factorial dengan Tail Recursive
func TailRecursive(result int, value int) int {
	if value <= 0 {
		return result
	} else {
		return TailRecursive(result*value, value-1)
	}
}

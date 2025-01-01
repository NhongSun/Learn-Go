package main

func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 0
	}
	return n * Factorial(n-1)
}

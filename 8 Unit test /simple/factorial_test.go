package main

import "testing"

func TestFactorial(t *testing.T) {
	// go test -run TestFactorial -v
	testCases := []struct {
		name     string
		n        int
		expected int
	}{
		{"Case 0", 0, 1},
		{"Case 2", 2, 2},
		{"Case 5", 5, 120},
		{"Case 10", 10, 3628800},
		{"Case 20", 20, 2432902008176640000},
		{"Case -1", -1, 0},
		{"Case -2", -2, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {})
		result := Factorial(tc.n)
		expectedResult := tc.expected

		if result != expectedResult {
			t.Errorf("Factorial(%d) = %d is wrong, the correct answer is %d", tc.n, result, expectedResult)
		}
	}
}

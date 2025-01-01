package main

import "testing"

func TestAdd(t *testing.T) {
	// add test before function name & _test.go to the file name to run the test
	// `go test` command will run all the test functions in the file
	// `go test -v` command will also show the output of the test
	// `go test -run TestAdd` command will only run TestAdd function name

	result := Add(2, 3)
	expectedResult := 5

	if result != expectedResult {
		t.Errorf("Add(2,3) = %d is wrong, the correct answer is %d", result, expectedResult)
	}

	testCases := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"Add positive numbers", 2, 3, 5},
		{"Add negative numbers", -1, -2, -3},
		{"Add zero", 0, 0, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {})
		result := Add(tc.a, tc.b)
		expectedResult := tc.expected

		if result != expectedResult {
			t.Errorf("Add(%d,%d) = %d is wrong, the correct answer is %d", tc.a, tc.b, result, expectedResult)
		}
	}
}

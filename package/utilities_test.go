package calculator

import "testing"

func TestSum(t *testing.T) {
	type testCases struct {
		name     string
		input1   int
		input2   int
		expected int
	}
	testCasesAll := []testCases{
		{
			name:     "correct output",
			input1:   5,
			input2:   10,
			expected: 15,
		},
		{
			name:     "incorrect output",
			input1:   5,
			input2:   11,
			expected: 17,
		},
	}

	for _, testCase := range testCasesAll {
		result := sum(testCase.input1, testCase.input2)
		if result != testCase.expected {
			t.Errorf("expected = %d, but got = %d", testCase.expected, result)
		}
	}
}

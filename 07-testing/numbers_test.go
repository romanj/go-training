package numbers

import (
	"testing"
)

func Test_Sum(t *testing.T) {
	testCases := []struct {
		name   string
		input  []int
		output int
	}{
		{
			name:   "testing with 1, 3, 5",
			input:  []int{1, 3, 5},
			output: 9,
		},
		{
			name:   "testing with 2, 4, 6",
			input:  []int{2, 4, 6},
			output: 12,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// we use the elipsis here to pass individual array elements
			result := Sum(tc.input...)
			if result != tc.output {
				t.Errorf("expected %d but got %d", tc.output, result)
			}
		})
	}
}

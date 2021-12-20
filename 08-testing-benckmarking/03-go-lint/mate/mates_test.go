package mate

import "testing"

type test struct {
	data     []int
	solution int
}

func TestSum(t *testing.T) {
	var tests = []test{
		{[]int{2, 4, 6}, 12},
		{[]int{1, 2, 3, 4}, 10},
		{[]int{1, 2, 3, 4, 5}, 15},
		{[]int{1, 2, 3, 4, 5, 6}, 21},
	}

	for _, element := range tests {
		var result = Sum(element.data...)
		if result != element.solution {
			t.Error("Expected", element.solution, "Got", result)
		}
	}
}

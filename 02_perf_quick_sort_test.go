package ca25l

import (
	"reflect"
	"testing"
)

func TestPerfQuickSort(t *testing.T) {
	var tests = []struct {
		input []int
		want  []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{3, 4, 1}, []int{1, 3, 4}},
		{[]int{8, 3, 10, 2, 7, 6, 9, 12}, []int{2, 3, 6, 7, 8, 9, 10, 12}},
	}
	for _, test := range tests {
		PerfQuickSort(test.input)
		if !reflect.DeepEqual(test.input, test.want) {
			t.Errorf("got %v, want %v", test.input, test.want)
		}
	}
}

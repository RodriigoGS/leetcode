package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	cases := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{
			[]int{2, 7, 11, 15},
			9,
			[]int{1, 0},
		},
		{
			[]int{3, 2, 4},
			6,
			[]int{2, 1},
		},
		{
			[]int{3, 3},
			6,
			[]int{1, 0},
		},
		{
			[]int{1, 2},
			4,
			nil,
		},
	}

	for _, c := range cases {
		r := TwoSum(c.nums, c.target)

		if !reflect.DeepEqual(r, c.expected) {
			t.Errorf("expect error %v. got %v", r, c.expected)
		}
	}
}

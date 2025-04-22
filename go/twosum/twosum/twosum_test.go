package twosum_test

import (
	"reflect"
	"testing"

	"github.com/vagnerclementino/algorithms/go/twosum/twosum"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{1, 2, 3, 4}, 7, []int{2, 3}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := twosum.Calculate(tt.nums, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum(%v, %d) = %v; want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}

package twosum_test

import (
	"reflect"
	"testing"

	"github.com/vagnerclementino/algorithms/go/twosum/twosum"
)

func TestTwoSum_Calculate(t *testing.T) {

	t.Parallel()
	tests := map[string]func(t *testing.T){
		"when receive a valid input with nums {2, 7, 11, 15} and target 9": func(t *testing.T) {

			t.Helper()
			// arrange
			nums := []int{2, 7, 11, 15}
			target := 9

			// act
			got := twosum.Calculate(nums, target)

			// assert
			if !reflect.DeepEqual(got, []int{0, 1}) {
				t.Errorf("Expected %v, but got %v", []int{0, 1}, got)
			}
		},
		"when receive a valid input with nums {3, 2, 4} and target 6": func(t *testing.T) {
			t.Helper()
			// arrange
			nums := []int{3, 2, 4}
			target := 6

			// act
			got := twosum.Calculate(nums, target)

			// assert
			if !reflect.DeepEqual(got, []int{1, 2}) {
				t.Errorf("Expected %v, but got %v", []int{1, 2}, got)
			}
		},
		"when receive a valid input with nums {3, 3} and target 6": func(t *testing.T) {
			t.Helper()
			// arrange
			nums := []int{3, 3}
			target := 6

			// act
			got := twosum.Calculate(nums, target)

			// assert
			if !reflect.DeepEqual(got, []int{0, 1}) {
				t.Errorf("Expected %v, but got %v", []int{0, 1}, got)
			}
		},
		"when receive a valid input with nums {1, 2, 3, 4} and target 7": func(t *testing.T) {
			t.Helper()
			// arrange
			nums := []int{1, 2, 3, 4}
			target := 7

			// act
			got := twosum.Calculate(nums, target)

			// assert
			if !reflect.DeepEqual(got, []int{2, 3}) {
				t.Errorf("Expected %v, but got %v", []int{2, 3}, got)
			}
		},
	}

	for name, run := range tests {
		run := run

		t.Run(name, func(t *testing.T) {
			t.Parallel()
			run(t)
		})
	}
}

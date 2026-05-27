package main

import (
	"fmt"
	"sort"
)

func verifyTriplets(triplets [][]int, target int) bool {
	for _, triplet := range triplets {
		sum := triplet[0] + triplet[1] + triplet[2]
		if sum != target {
			return false
		}
	}
	return true
}
func threeSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	triplets := [][]int{}
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		low, high := i+1, len(nums)-1
		for low < high {
			sum := nums[low] + nums[high] + nums[i]
			if sum == target {
				triplets = append(triplets, []int{nums[low], nums[high], nums[i]})
				for low < high && nums[low] == nums[i+1] {
					low++
				}
				for low < high && nums[high] == nums[high-1] {
					high++
				}
				low++
				high--
			} else if sum > target {
				high--
			} else {
				low++
			}
		}
	}
	return triplets
}

func main() {
	nums := []int{
		-10, -7, -3, -2, -1, 0, 0, 1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		-5, -6, -4, 11, 13, -8, -9, 12, -11, 15, -12, 14, -13, 20, -20,
		17, -17, 16, -16, 18, -18, 19, -19, 21, -21,
	}
	target := 10
	triplets := (threeSum(nums, target))
	fmt.Println("Triplets found:")
	for _, t := range triplets {
		fmt.Println(t)
	}
	if verifyTriplets(triplets, target) {
		fmt.Println("✅ All triplets are valid and sum to target.")
	} else {
		fmt.Println("❌ Invalid triplet(s) found.")
	}
}

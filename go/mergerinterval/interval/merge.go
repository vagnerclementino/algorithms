package interval

import (
	"cmp"
	"slices"
)

func merge(intervals [][]int) [][]int {
	slices.SortFunc(intervals, func(a, b []int) int {
		return cmp.Compare(a[0], b[0])
	})
	merged := [][]int{}
	for _, interval := range intervals {
		lastMerged := len(merged) - 1
		if len(merged) == 0 || merged[lastMerged][1] < interval[0] {
			merged = append(merged, interval)
		} else {
			merged[lastMerged][1] = max(merged[lastMerged][1], interval[1])
		}
	}
	return merged
}

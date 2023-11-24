package main

import (
	"fmt"
	"math"
)

func main() {
	t := smallestDifference([]int{-1, 5, 10, 20, 28, 3}, []int{26, 134, 135, 15, 17})
	fmt.Println(t)

}
func smallestDifference(array1, array2 []int) []int {

	smallTuple := []int{0, 0}
	quickSort(array1, 0, len(array1)-1)
	quickSort(array2, 0, len(array2)-1)
	i, j := 0, 0
	smallDiff := math.MaxInt
	for {
		if i == len(array1) || j == len(array2) {
			break
		}
		currDiff := absDiff(array1[i], array2[j])
		if currDiff < smallDiff {
			smallTuple[0] = array1[i]
			smallTuple[1] = array2[j]
			smallDiff = currDiff
		}
		if smallDiff == 0 {
			break
		}

		if array1[i] <= array2[j] {
			i++
		} else {
			j++
		}
	}
	return smallTuple
}

func absDiff(m, n int) int {
	var diff float64
	if (m < 0 && n < 0) || (m > 0 && n > 0) {
		diff = math.Abs(float64(n)) - math.Abs(float64(m))
	} else {
		diff = math.Abs(float64(n)) + math.Abs(float64(m))
	}
	return int(math.Abs(diff))
}

func quickSort(arr []int, low int, high int) {
	if low < high {
		pi := partion(arr, low, high)
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func partion(arr []int, low int, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

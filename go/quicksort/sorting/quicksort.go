package sorting

import "fmt"

func QuickSort(a []int, low, high int) {
	if low < high {
		p := partition(a, low, high)
		fmt.Println(p)
		QuickSort(a, low, p-1)
		QuickSort(a, p+1, high)
	}

}

func partition(a []int, low, high int) int {
	//select the last element as pivot
	pivot := a[high] 
	pivotIdx := low
	//partition the array in one linear scan for the pivot
	for j := low; j < high; j++ {
		if a[j] < pivot {
			a[pivotIdx], a[j] = a[j], a[pivotIdx]
			pivotIdx++
		}
	}
	// After the looping, the pivotIdx has the right position for the pivot
	// in the array
	a[pivotIdx], a[high] = a[high], a[pivotIdx]
	return pivotIdx
}
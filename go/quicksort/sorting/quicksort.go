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
	pivot := a[high]
	pivotIdx := low
	for j := low; j < high; j++ {
		if a[j] < pivot {
			a[pivotIdx], a[j] = a[j], a[pivotIdx]
			pivotIdx++
		}
	}
	a[pivotIdx], a[high] = a[high], a[pivotIdx]
	return pivotIdx
}

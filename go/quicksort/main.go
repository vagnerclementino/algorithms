package main

import (
	"fmt"

	"github.com/vagnerclementino/algorithm/go/quicksort/go/quicksort/sorting"
)

func main() {
	a := []int{5, 6, 7, 2, 1, 0}
	sorting.QuickSort(a, 0, len(a)-1)
	fmt.Println(a)
}

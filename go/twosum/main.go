package main

import (
	"fmt"

	"github.com/vagnerclementino/algorithms/go/twosum/twosum"
)

func main() {
	nums := []int{-3, 4, 3, 90}
	target := 0
	fmt.Println(twosum.Calculate(nums, target))
}

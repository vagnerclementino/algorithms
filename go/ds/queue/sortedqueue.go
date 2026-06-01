package ds

import (
	"fmt"
	"slices"
)

type SortedQueue struct {
	values []int
}

func NewSortedQueue() *SortedQueue {
	return &SortedQueue{values: []int{}}
}

func (q *SortedQueue) Add(item int) {
	for idx, value := range q.values {
		if item > value {
			q.values = slices.Insert(q.values, idx, item)
			return
		}
	}

	q.values = append(q.values, item)
}
func (q *SortedQueue) String() string {
	output := "-start-> "
	for _, v := range q.values {
		output += fmt.Sprintf("%d ", v)
	}
	return output
}

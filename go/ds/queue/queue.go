package ds

import (
	"fmt"
)

type Queue struct {
	values []int
}

func NewQueue() *Queue {
	return &Queue{values: []int{}}
}

func (q *Queue) Enqueue(item int) {
	q.values = append(q.values, item)
}

func (q *Queue) Dequeue() (error, int) {
	if q == nil || len(q.values) == 0 {
		return fmt.Errorf("the queue is empty"), 0
	}
	value := q.values[0]
	q.values = q.values[1:]
	return nil, value
}

func (q *Queue) String() string {
	output := "-start-> "
	for _, v := range q.values {
		output += fmt.Sprintf("%d ", v)
	}
	return output
}

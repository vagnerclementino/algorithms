package ds

import "fmt"

type Stack struct {
	values []int
}

func NewStack() *Stack {
	return &Stack{values: []int{}}
}

func (s *Stack) Push(item int) {
	s.values = append(s.values, item)
}

func (s *Stack) Pop() int {
	value := s.values[len(s.values)-1]
	if len(s.values) > 0 {
		s.values = s.values[:len(s.values)-1]
	}
	return value
}

func (s *Stack) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *Stack) String() string {
	output := "-start-> "
	for i := len(s.values) - 1; i >= 0; i-- {
		output += fmt.Sprintf("%d ", s.values[i])
	}
	return output
}

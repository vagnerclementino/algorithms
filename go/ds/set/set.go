package ds

import "fmt"

type Set struct {
	values map[int]struct{}
}

func NewSet() *Set {
	return &Set{values: map[int]struct{}{}}
}

func (s *Set) Add(item int) {
	s.values[item] = struct{}{}
}

func (s *Set) Delete(item int) {
	delete(s.values, item)
}

func (s *Set) Contains(item int) bool {
	if _, ok := s.values[item]; ok {
		return true
	} else {
		return false
	}
}

func (s *Set) Intersect(other *Set) *Set {
	intesect := NewSet()
	for k, _ := range s.values {
		if contains := other.Contains(k); contains {
			intesect.Add(k)
		}
	}
	return intesect
}

func (s *Set) Values() []int {
	values := make([]int, 0, len(s.values))
	for k, _ := range s.values {
		values = append(values, k)
	}
	return values
}

func (s *Set) Union(other *Set) *Set {
	unionSet := NewSet()

	for k, _ := range s.values {
		unionSet.Add(k)
	}

	for v := range other.Values() {
		unionSet.Add(v)
	}
	return unionSet
}

func (s *Set) String() string {
	output := "("
	for k, _ := range s.values {
		output += fmt.Sprintf("%d,", k)
	}
	return output[:len(output)-1] + ")"
}

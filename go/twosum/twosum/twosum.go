package twosum

func Calculate(nums []int, target int) []int {
	m := map[int]int{}

	for i, n := range nums {
		complement := target - n
		if j, ok := m[complement]; ok {
			return []int{j, i}
		}
		m[n] = i
	}

	return []int{}
}

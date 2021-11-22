package array

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		if _, ok := m[complement]; ok {
			return []int{i, m[complement]}
		}
		m[nums[i]] = i

	}

	return nil

}

package array

func TwoSums(nums []int, target int) []int {
	h := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if j, ok := h[target-nums[i]]; ok {
			return []int{i, j}
		}
		h[nums[i]] = i
	}

	return nil
}

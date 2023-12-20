package backtracking

// 找出所有递归子序列
func findSubsequences(nums []int) [][]int {
	path := make([]int, 0, len(nums))
	res := make([][]int, 0)
	dfsSubsequence(nums, 0, path, &res)
	return res
}

func dfsSubsequence(nums []int, start int, path []int, res *[][]int) {
	// 回溯终止条件
	if len(path) >= 2 {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
	}

	used := make(map[int]bool) // 用于对同层元素去重
	for i := start; i < len(nums); i++ {
		if used[nums[i]] {
			continue
		}

		if len(path) == 0 || nums[i] >= path[len(path)-1] {
			path = append(path, nums[i])
			used[nums[i]] = true
			dfsSubsequence(nums, i+1, path, res)
			path = path[:len(path)-1]
		}
	}
}

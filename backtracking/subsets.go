package backtracking

// 给定一个整数数字nums, 返回所有的子集
func subsets(nums []int) [][]int {
	path := make([]int, 0, len(nums))
	res := make([][]int, 0)
	backtrackSubsets(nums, 0, path, &res)
	return res
}

func backtrackSubsets(nums []int, start int, path []int, res *[][]int) {
	// 收集子集
	tmp := make([]int, len(path))
	copy(tmp, path)
	*res = append(*res, tmp)

	for i := start; i < len(nums); i++ {
		path = append(path, nums[i])
		backtrackSubsets(nums, i+1, path, res)
		path = path[:len(path)-1]
	}
}

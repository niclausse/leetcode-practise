package backtracking

import "sort"

// 组合总和
// 给定一个无重复元素的整数数组candidates, 和一个目标值target
// 找出candidates中所有可以使数字和为target的组合 (元素可被重复使用)
// 示例:
// 输入：candidates = [2,3,6,7], target = 7,
// 所求解集为： [ [7], [2,2,3] ]
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := make([][]int, 0)
	path := make([]int, 0, len(candidates))
	backtrackSum(candidates, target, 0, path, &res)
	return res
}

func backtrackSum(candidates []int, target, start int, path []int, res *[][]int) {
	if target == 0 {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}

	for i := start; i < len(candidates); i++ {
		if candidates[i] > target {
			break
		}

		path = append(path, candidates[i])
		backtrackSum(candidates, target-candidates[i], i, path, res)
		path = path[:len(path)-1]
	}
}

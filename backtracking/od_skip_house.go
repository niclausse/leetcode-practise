package backtracking

import "math"

// 跳房子, count为房子的总格数,  steps为每次可以跳的格子数
// 返回三个回合可以跳到房子最后的步数组合 (只返回索引和最小的组合)
func skipHouse(count int, steps []int) []int {
	path := make([]int, 0, 3)
	res := make([]int, 3)
	minIndexSum := math.MaxInt
	dfsSkipHouse(steps, count, 0, 0, path, &res, &minIndexSum)
	return res
}

// 回溯找到
func dfsSkipHouse(steps []int, leftCount, start int, indexSum int, path []int, res *[]int, minIndexSum *int) {
	// 回溯终止条件
	if len(path) == 3 && leftCount == 0 && indexSum < *minIndexSum {
		*minIndexSum = indexSum
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = tmp
		return
	}

	for i := start; i < len(steps); i++ {
		if steps[i] > leftCount {
			continue
		}

		path = append(path, steps[i])
		dfsSkipHouse(steps, leftCount-steps[i], i+1, indexSum+i, path, res, minIndexSum)
		path = path[:len(path)-1]
	}
}

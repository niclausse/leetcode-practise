package backtracking

// 组合总数 找出所有和为n的k个数的组合,  组合元素必须为1~9的数字&&元素不能重复
func combinationSum3(n, k int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0, k)
	dfsSum3(n, k, 1, path, &res)
	return res
}

func dfsSum3(n, k, start int, path []int, res *[][]int) {
	var sum int
	for _, v := range path {
		sum += v
	}
	if sum == n && len(path) == k {
		tmp := make([]int, k)
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}

	for i := start; i <= 9; i++ { // 横向遍历
		if sum+i > n || len(path)+1 > k {
			break
		}

		path = append(path, i)
		dfsSum3(n, k, i+1, path, res) // 纵向递归
		path = path[:len(path)-1]     // 撤销本次回溯结果,  开始下次回溯
	}
}

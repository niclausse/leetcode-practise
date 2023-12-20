package backtracking

// 组合问题 给定两个整数n k, 返回1-n范围内所有可能的k的个数的组合
func combine(n, k int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0, k)
	dfs(n, k, 1, path, &res)
	return res
}

func dfs(n, k, start int, path []int, res *[][]int) {
	if len(path) == k {
		tmp := make([]int, k)
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}

	for i := start; i <= n; i++ { // 横向遍历
		if n-i+1 < k-len(path) { // 剪枝 / 剩余元素不够
			break
		}
		path = append(path, i)
		dfs(n, k, i+1, path, res) // 纵向递归
		path = path[:len(path)-1] // 撤销本次回溯结果, 开始下次回溯
	}
}

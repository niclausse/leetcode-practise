package backtracking

// 分割回文串: 给定一个字符串s, 将s分割成一些子串, 使得每个子串都是回文串, 返回所有可能的分割方案
func partition(s string) [][]string {
	res := make([][]string, 0)
	path := make([]string, 0, len(s))
	dfsPartition(s, 0, path, &res)
	return res
}

func dfsPartition(s string, start int, path []string, res *[][]string) {
	// 回溯终止条件
	if start == len(s) {
		tmp := make([]string, len(path))
		copy(tmp, path) // path还要在后面的回溯中继续使用,  元素会发生变化
		*res = append(*res, tmp)
		return
	}

	for i := start; i < len(s); i++ {
		str := s[start : i+1]
		if isPalindrome(str) {
			path = append(path, str)
			dfsPartition(s, i+1, path, res)
			path = path[:len(path)-1]
		}
	}
}

func isPalindrome(s string) bool {
	if s == "" {
		return false
	}
	lft, rgt := 0, len(s)-1
	for lft < rgt {
		if s[lft] != s[rgt] {
			return false
		}
		lft++
		rgt--
	}
	return true
}

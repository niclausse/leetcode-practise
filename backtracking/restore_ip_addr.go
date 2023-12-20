package backtracking

import (
	"strconv"
	"strings"
)

// 给定一个只包含数字的字符串s, 返回所有可能的IP地址
func restoreIpAddresses(s string) []string {
	res := make([]string, 0)
	path := make([]string, 0, 4)
	backtrackIP(s, 0, path, &res)
	return res
}

func backtrackIP(s string, start int, path []string, res *[]string) {
	if len(path) == 4 && start == len(s) { // ip4段
		*res = append(*res, strings.Join(path, "."))
		return
	}

	for i := start; i < len(s); i++ {
		if i != start && s[start] == '0' { // 前导0  无效
			break
		}

		str := s[start : i+1]
		num, _ := strconv.Atoi(str)
		if num < 0 || num > 255 { // 数字不在0~255 无效
			break
		}

		path = append(path, str)
		backtrackIP(s, i+1, path, res)
		path = path[:len(path)-1]
	}
}

package backtracking

var digit2letters = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

// 电话号码的字母组合 - 输入一个仅包含数字2-9的字符串, 返回所有它能表示的字母组合
func letterCombinations(digits string) []string {
	candidates := make([]string, 0)
	for i := 0; i < len(digits); i++ {
		candidates = append(candidates, digit2letters[digits[i]])
	}

	res := make([]string, 0)
	var path string
	backtracking(candidates, 0, path, &res)
	return res
}

func backtracking(candidates []string, start int, path string, res *[]string) {
	if start == len(candidates) { // candidates遍历完成时结束
		*res = append(*res, path)
		return
	}

	str := candidates[start] // 当前回溯的字母组合
	for i := 0; i < len(str); i++ {
		path += str[i : i+1]
		backtracking(candidates, start+1, path, res) // 回溯下一个字母组合
		path = path[:len(path)-1]
	}
}

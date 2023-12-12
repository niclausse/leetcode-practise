package stack

// GenerateParenthesis 括号生成, n表示提供n对'(' ')'
// 要求返回所有合理的括号组合
func GenerateParenthesis(n int) []string {
	var res []string
	generate(&res, "", n, n)
	return res
}

// right, left剩余的左右括号数
func generate(res *[]string, str string, right, left int) {
	if right == 0 && left == 0 {
		*res = append(*res, str)
		return
	}

	if right == left { // 剩余的左右括号数相等,  则下一个只能用左括号
		str += "("
	} else if left < right { // 剩余的左括号数小于剩余的右括号数, 下一个既可以用左括号 也可以用右括号
		if left > 0 {
			generate(res, str+"(", right, left-1)
		}
		generate(res, str+")", right-1, left)
	}
}

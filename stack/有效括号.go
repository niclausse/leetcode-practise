package stack

// 题20
// 左括号压入栈中， 扫描到右括号时，从栈顶取出一个括号进行匹配
func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	var stack []rune
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		}

		if len(stack) == 0 {
			return false
		}

		if c == ')' {
			popper := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if popper != '(' {
				return false
			}
		}

		if c == ']' {
			popper := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if popper != '[' {
				return false
			}
		}

		if c == '}' {
			popper := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if popper != '{' {
				return false
			}
		}
	}

	if len(stack) == 0 {
		return true
	}

	return false
}

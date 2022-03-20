package string

// 题号20
// 验证字符串是否为回文串， 只考虑字母和数字
// 思路：双指针判断
// 空间复杂度：O(1)
// 时间复杂度：O(|s|)
func isPalindrome(s string) bool {
	if s == "" {
		return true
	}

	lft, rgt := 0, len(s)-1

	for lft < rgt {
		for lft < rgt && !valid(rune(s[lft])) {
			lft++
		}

		for lft < rgt && !valid(rune(s[rgt])) {
			rgt--
		}

		if s[lft] != s[rgt] {
			return false
		}

		lft++
		rgt--
	}

	return true
}

func valid(c rune) bool {
	if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
		return true
	}
	return false
}

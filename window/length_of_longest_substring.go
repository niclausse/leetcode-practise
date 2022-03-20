package window

// 题3：无重复字符的最大子串的长度
func lengthOfLongestSubString(s string) int {
	if len(s) == 0 {
		return 0
	}

	exist := make(map[byte]bool)
	result, start, end := 1, 0, 0

	for end < len(s) {
		for start < end && exist[s[end]] {
			delete(exist, s[start])
			start++
		}

		result = max(result, end-start+1)
		exist[s[end]] = true
		end++
	}

	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

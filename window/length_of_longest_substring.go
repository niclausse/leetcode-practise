package window

func findLongestSubString(s string) []string {
	hash := make(map[byte]bool)
	subHash := make(map[int][]string)
	length, lft, rgt := 0, 0, 1

	for lft < rgt && rgt < len(s) {
		for hash[s[rgt]] && lft < rgt {
			delete(hash, s[lft])
			lft++
		}

		rgt++
		hash[s[rgt]] = true
		length = max(length, rgt-lft+1)
		subHash[length] = append(subHash[length], s[lft:rgt+1])
	}

	return subHash[length]
}

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

// 思路： 使用双指针, 右指针不断向后移动， 判断区间内有无重复字符
func lengthOfLongestSubStringPractise2(s string) int {
	if len(s) == 0 {
		return 0
	}

	result, rgt, lft := 1, 0, 0

	seen := make(map[byte]bool)
	for rgt < len(s) {
		// 处理左指针需要后移的case
		for lft < rgt && seen[s[rgt]] {
			delete(seen, s[lft])
			lft++
		}

		// 处理右指针后移的case
		tmpLen := rgt - lft + 1
		if result < tmpLen {
			result = tmpLen
		}

		seen[s[rgt]] = true
		rgt++
	}

	return result
}

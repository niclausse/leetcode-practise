package string

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		prefix = lcp(prefix, strs[i])
		if prefix == "" {
			return prefix
		}
	}

	return prefix
}

func lcp(str1, str2 string) string {
	idx := 0
	for idx < len(str1) && idx < len(str2) && str1[idx] == str2[idx] {
		idx++
	}
	return str1[:idx]
}

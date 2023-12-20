package string

// 最长回文串 - 双指针中心拓展法 cbbd
func longestPalindrome(s string) string {
	var left, right int
	for i := 0; i < len(s); i++ {
		if lft, rgt := extend(s, i, i); rgt-lft > right-left {
			left = lft
			right = rgt
		}
		if lft, rgt := extend(s, i, i+1); rgt-lft > right-left {
			left = lft
			right = rgt
		}
	}
	return s[left : right+1]
}

// 以i j 为中心, 寻找最长回文串, 返回左右下标
func extend(s string, i, j int) (int, int) {
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
	}
	return i + 1, j - 1
}

// 最长回文串 - 动态规划法
func longestPalindromeDP(s string) string {
	// 1' 定义dp数组 dp数组 dp[i][j]表示字符串s在[i:j]范围内是否为回文串
	// 2' dp数组递推公式
	// 如果s[i] != s[j], dp[i][j] = false
	// 如果s[i] == s[j],
	// 情况1, i == j, dp[i][j] = true
	// 情况2, 下标i,j相差1, 如aa, 则dp[i][j] = true
	// 情况3, 当i,j相差大于1的时候, 取决于子区间dp[i+1][j-1]

	maxLength, left, right := 0, 0, 0
	dp := make([][]bool, len(s))

	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					dp[i][j] = true
				} else if dp[i+1][j-1] {
					dp[i][j] = true
				}
			}

			if dp[i][j] && j-i+1 > maxLength {
				maxLength = j - i + 1
				left = i
				right = j
			}
		}
	}

	return s[left : right+1]
}

// 最长回文子串 - 动态规划
func longestPalindromeLength(s string) int {
	// dp数组 dp[i][j]表示字符串s在[i:j]范围内最长的回文子串的长度
	// 对于回文串
	// 如果s[i] == s[j], dp[i][j] = dp[i+1][j-1] + 2
	// 如果s[i] != s[j], dp[i][j] = max(dp[i+1][j], dp[i][j-1])

	max := func(x, y int) int {
		if x >= y {
			return x
		}
		return y
	}

	dp := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(s))
		dp[i][i] = 1
	}

	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}

	return dp[0][len(s)-1]
}

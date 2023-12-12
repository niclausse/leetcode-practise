package dp

// 给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
// 计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
// 可以认为每种硬币的数量是无限的。

// dp[amount]=min{dp[amount-coins[0],dp[amount-coins[1],...,dp[amount-coins[n-1]}+1
func leastCoins(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = amount + 1 // 先将dp数组设置一个最大值
	}

	for i := 1; i <= amount; i++ { // 总额
		for j := 0; j < len(coins); j++ {
			if i >= coins[j] {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

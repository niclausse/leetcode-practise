package dp

// 01背包问题
// n件物品 第i件物品的重量为weight[i], 价值为value[i]
// 背包有最大承受质量
// 每件物品只能用一次, 求解背包能装物品的最大价值

func BagMaxValue(weights, values []int, bagWeight int) int {
	// dp[j] = max(dp[j], dp[i - weight[i]] + value[i])
	dp := make([]int, bagWeight+1)

	for i := 0; i < len(weights); i++ { // 遍历物品
		for j := bagWeight; j >= weights[i]; j-- {
			dp[j] = max(dp[j], dp[j-weights[i]]+values[i])
		}
	}

	return dp[bagWeight]
}

package dp

// 完全背包问题
// n件物品 第i件物品的重量为weight[i], 价值为value[i]
// 背包有最大承受质量
// 每种物品可以无限制用, 求解背包能装物品的最大价值

func MaxBagValueFull(weights, values []int, maxBagWeight int) int {
	dp := make([]int, maxBagWeight+1)

	// 遍历物品
	for i := 0; i < len(weights); i++ {
		// 遍历背包
		for j := weights[i]; j <= maxBagWeight; j++ {
			dp[j] = max(dp[j], dp[j-weights[i]]+values[i])
		}
	}

	return dp[maxBagWeight]
}

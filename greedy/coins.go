package greedy

import "sort"

// 给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
// 计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
// 可以认为每种硬币的数量是无限的。
func leastCoins(coins []int, amount int) int {
	if len(coins) == 0 {
		return -1
	}

	sort.Ints(coins)

	res := 0
	for i := len(coins) - 1; i >= 0; i-- {
		for amount > coins[i] {
			res++
			amount -= coins[i]
		}
	}

	if amount != 0 {
		return -1
	}

	return res
}

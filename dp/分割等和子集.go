package dp

//给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
//
//示例 1：
//
//输入：nums = [1,5,11,5]
//输出：true
//解释：数组可以分割成 [1, 5, 5] 和 [11] 。
//示例 2：
//
//输入：nums = [1,2,3,5]
//输出：false
//解释：数组不能分割成两个元素和相等的子集。
//
//
//提示：
//
//1 <= nums.length <= 200
//1 <= nums[i] <= 100

// 转化为01背包问题：背包容量为sum/2， 所装物品的最大价值必须等于sum/2
func canPartition(nums []int) bool {
	var sum int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	if sum%2 != 0 {
		return false
	}

	target := sum / 2

	dp := make([]int, 10001)
	// 遍历物品
	for i := 0; i < len(nums); i++ {
		// 遍历背包
		for j := target; j >= nums[i]; j-- {
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}

	return dp[target] == target
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

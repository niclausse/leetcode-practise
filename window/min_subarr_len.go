/*
Package window includes solutions about window
*/
package window

import "math"

// 题209：长度最小的子数组
func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	result := math.MaxInt

	start, end, sum := 0, 0, 0

	for end < len(nums) {
		sum += nums[end]
		for sum >= target {
			result = min(result, end-start+1)
			sum -= nums[start]
			start++
		}
		end++
	}

	if result == math.MaxInt {
		return 0
	}

	return result
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

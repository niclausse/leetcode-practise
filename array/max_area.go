package array

import "github.com/penglin1995/leetcode-practise/tools"

// 题11：盛最多水的容器

func maxArea(height []int) (area int) {

	lft, rgt := 0, len(height)-1

	for lft < rgt {
		area = tools.Max(tools.Min(height[lft], height[rgt])*(rgt-lft), area)

		if height[lft] <= height[rgt] {
			lft++
		} else {
			rgt--
		}
	}

	return
}

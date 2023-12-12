package leetcode_practise

import (
	"sort"
	"strconv"
)

func maxNum(nums []string, n int) int {
	sort.Slice(nums, func(i, j int) bool {
		ni, _ := strconv.Atoi(nums[i])
		nj, _ := strconv.Atoi(nums[j])
		return ni < nj
	})

	ns := strconv.Itoa(n)
	// 位数相同, 从高位到低位， 选小于等于的数字

	for i := 0; i < len(ns); i++ {
		e, _ := strconv.Atoi(ns[i : i+1])

	}
}

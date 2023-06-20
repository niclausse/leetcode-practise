package stack

// nums1 = [4,1,2], nums2 = [1,3,4,2]
func nextGreaterElement1(nums1, nums2 []int) []int {
	var (
		stack []int
		res   = make([]int, 0, len(nums1))
		mp    = make(map[int]int)
	)

	for i := len(nums2) - 1; i >= 0; i-- {
		// 如果当前元素大于或等于栈顶元素， 则将栈顶元素出栈
		for len(stack) > 0 && stack[len(stack)-1] <= nums2[i] {
			stack = stack[:len(stack)-1]
		}

		// 记录当前元素与下一个更大元素的映射关系
		if len(stack) == 0 {
			mp[nums2[i]] = -1
		} else {
			mp[nums2[i]] = stack[len(stack)-1]
		}

		stack = append(stack, nums2[i])
	}

	for _, v := range nums1 {
		res = append(res, mp[v])
	}

	return res
}

// 循环数组 https://leetcode.cn/problems/next-greater-element-ii/
//func nextGreaterElements2(nums []int) []int {
//	var stack []int
//	for i := 0; i < len(nums); i++ {
//		for len(stack) > 0 &&
//	}
//}

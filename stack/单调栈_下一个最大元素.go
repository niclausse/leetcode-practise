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

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums1))
	table := make(map[int]int, len(nums1)) // 元素在nums中对应的下标
	for i, v := range nums1 {
		res[i] = -1 // 设置默认值
		table[v] = i
	}

	var stack []int

	// 使用单调栈寻找nums中下一个更大选素
	for i := 0; i < len(nums2); i++ {
		j, ok := table[nums2[i]]
		if !ok {
			continue
		}

		for len(stack) > 0 && nums2[i] > stack[len(stack)-1] {
			// 出栈
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[j] = top
		}

		// 入栈
		stack = append(stack, nums2[i])
	}

	return res
}

// 给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，其中 answer[i] 是指对于第 i 天，下一个更高温度出现在几天后。如果气温在这之后都不会升高，请在该位置用 0 来代替。
func dailyTemperatures(temperatures []int) []int {
	var (
		res   = make([]int, len(temperatures))
		stack []int // 记录下标
	)

	for i := 0; i < len(temperatures); i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			// 出栈
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			res[top] = i - top
		}

		// 入栈
		stack = append(stack, i)
	}

	return res
}

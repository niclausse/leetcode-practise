package array

// 删除有序数组nums中的重复元素, 返回删除后的数组长度
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 双指针
	// 快指针用于遍历数组
	// 慢指针用于接收下一个非重复元素
	slow := 1
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast] // 将当前非重复元素放到慢指针指向的位置
			slow++
		}
	}

	return slow
}

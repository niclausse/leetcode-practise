package array

// 题283： 移动零
// 给定一个数组nums， 将数组中的所有0移动到数组末尾， 并保持非0元素的相对位置

// 思路： 双指针
// 左指针指向已处理好的序列的末端， 右指针指向未处理好的序列的头部
func moveZeros(nums []int) {
	lft, rgt := 0, 0

	for rgt < len(nums) {
		if nums[rgt] != 0 {
			nums[lft], nums[rgt] = nums[rgt], nums[lft]
			lft++
		}

		rgt++
	}
}

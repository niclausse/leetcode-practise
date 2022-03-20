package leetcode_practise

func binarySearch(nums []int, target int) int {
	lft, rgt := 0, len(nums)-1

	for lft < rgt {
		mid := lft + (rgt-lft)/2

		midValue := nums[mid]

		if midValue == target {
			return mid
		} else if midValue > target {
			rgt = mid - 1
		} else {
			lft = lft + 1
		}
	}

	return -1
}

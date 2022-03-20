package golang

import "fmt"

func increaseCapacity(nums []int) []int {
	fmt.Printf("old.len: %d\n", len(nums))
	fmt.Printf("old.cap: %d\n", cap(nums))

	nums = append(nums, 3, 4, 5)
	fmt.Printf("new.len: %d\n", len(nums))
	fmt.Printf("new.cap: %d\n", cap(nums))

	return nums
}

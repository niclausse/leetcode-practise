package golang

import (
	"fmt"
	"testing"
)

func TestIncreaseSlice(t *testing.T) {
	nums := []int{1, 2}
	arr := increaseCapacity(nums)

	fmt.Println(nums)
	fmt.Println(arr)
}

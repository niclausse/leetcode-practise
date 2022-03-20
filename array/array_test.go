package array

import (
	"fmt"
	"testing"
)

func TestMoveZeros(t *testing.T) {
	nums := []int{1, 0, 2, 3, 0, 6, 7, 0}
	moveZeros(nums)
	fmt.Println(nums)
}

func TestMaxArea(t *testing.T) {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}

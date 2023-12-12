package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{9, 5, 3, 4, 0, 1, 6}
	BubbleSort(arr)
	fmt.Println(arr)
}

func TestQuickSort(t *testing.T) {
	arr := []int{9, 5, 3, 4, 0, 1, 6}
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

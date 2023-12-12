package search

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{0, 1, 2, 3, 5, 6, 7, 8, 9},
	}

	r1 := BinarySearch(tests[0], 7)
	if r1 != 6 {
		t.Errorf("7存在于tests[0]下标6位置")
		return
	}

	r2 := BinarySearch(tests[1], 4)
	if r2 != -1 {
		t.Errorf("4不存在于tests[1]")
		return
	}
}

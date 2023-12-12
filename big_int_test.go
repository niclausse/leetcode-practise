package leetcode_practise

import (
	"strconv"
	"testing"
)

func TestAddBigint(t *testing.T) {
	res := addBigInt("123", "1")
	r, _ := strconv.Atoi(res)
	if r != 124 {
		t.Errorf("res shouldbe 124, but got %s", res)
	}
}

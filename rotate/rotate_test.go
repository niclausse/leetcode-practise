package rotate

import (
	"fmt"
	"testing"
)

func TestRotate(t *testing.T) {
	in := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	rotate(in)
	fmt.Println(in)
}

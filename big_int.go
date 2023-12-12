package leetcode_practise

import "strconv"

func addBigInt(a, b string) string {
	var res string
	quota := 0 // 进位
	for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0; {
		var x, y int
		if i >= 0 {
			x, _ = strconv.Atoi(a[i : i+1])
		} else {
			x = 0
		}
		if j >= 0 {
			y, _ = strconv.Atoi(b[j : j+1])
		} else {
			y = 0
		}

		s := x + y + quota
		quota = s / 10
		v := s % 10
		res = strconv.FormatInt(int64(v), 10) + res

		i--
		j--
	}

	return res
}

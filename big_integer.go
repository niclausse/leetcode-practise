package leetcode_practise

import "strconv"

func add(a, b string) (string, error) {
	lenA := len(a)
	lenB := len(b)

	max := _max(lenA, lenB)

	if lenA < max {
		prefix := ""
		for i := 0; i < max-lenA; i++ {
			prefix += "0"
		}

		a = prefix + a
	}

	if lenB < max {
		prefix := ""
		for i := 0; i < max-lenB; i++ {
			prefix += "0"
		}

		b = prefix + b
	}

	var quota int64 // 进位
	var result string

	for i := max - 1; i >= 0; i-- {
		tmpA, err := strconv.ParseInt(string(a[i]), 10, 64)
		if err != nil {
			return "", err
		}

		tmpB, err := strconv.ParseInt(string(b[i]), 10, 64)
		if err != nil {
			return "", err
		}

		var tmp int64
		sum := tmpA + tmpB + quota
		if sum > 10 && i != 0 {
			tmp = sum - 10
			quota = 1
		} else {
			tmp = sum
			quota = 0
		}

		result = strconv.FormatInt(tmp, 10) + result
	}

	return result, nil
}

func _max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

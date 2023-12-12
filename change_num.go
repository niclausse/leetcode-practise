package leetcode_practise

import (
	"strconv"
)

func changeNum(num int) int {

	var arr []int
	ws := 10
	for num/ws != 0 {
		arr = append(arr, (num%ws)/(ws/10))
		ws *= 10
	}

	arr = append(arr, (num%ws)/(ws/10))

	tmpData := 0
	for i := len(arr) - 1; i >= 0; i-- {
		cur := arr[i]
		max := arr[0]

		wz := 0

		for j := 0; j < i; j++ {
			if arr[j] > max {
				max = arr[j]
				wz = j
			}
		}
		if max > cur {
			tmpData = arr[wz]

			arr[wz] = arr[i]
			arr[i] = tmpData
			break
		}
	}

	result := 0
	step := 1
	for i := 0; i < len(arr); i++ {
		result += arr[i] * step
		step *= 10
	}

	return result
}

// 输入一个数字，交换其中的两位， 返回交换后的最大值
// 1‘ 暴力解法
func change1(num int) (result int) {
	result = num

	s := []byte(strconv.Itoa(num))
	for i := range s {
		for j := range s[:i] {
			s[i], s[j] = s[j], s[i]
			tmp, _ := strconv.Atoi(string(s))
			if tmp > result {
				result = tmp
			}
			s[j], s[i] = s[i], s[j]
		}
	}

	return result
}

// 2' 贪心算法
// 右边越大的数字与左边越小的数字交换， 才能使交换后的数字越大
// 因此， 尝试将右边较大数字与左边较小数字进行交换
func change2(num int) int {
	s := []byte(strconv.Itoa(num))
	n := len(s)

	maxIdx, idx1, idx2 := n-1, -1, -1

	for i := n - 1; i >= 0; i-- {
		if s[i] > s[maxIdx] {
			maxIdx = i
		} else {
			idx1, idx2 = i, maxIdx
		}
	}

	s[idx1], s[idx2] = s[idx2], s[idx1]
	result, _ := strconv.Atoi(string(s))
	return result
}

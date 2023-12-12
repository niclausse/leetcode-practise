package greedy

import "strconv"

// 输入一个数字，交换其中的两位， 返回交换后的最大值
// 最优解问题
// 越靠右的较大值与越靠左的较小值交换 ==> 一贪贪最大，二贪贪最左
func change(num int) int {
	s := []byte(strconv.Itoa(num))

	maxIdx := len(s) - 1 // 初始假定最右侧的数字最大（贪最大+贪最右）
	idx1, idx2 := -1, -1

	// 从右往左遍历， 左侧有更小的数字， 则可以进行局部交换（局部最优）
	for i := range s {
		if s[i] < s[maxIdx] {
			idx1, idx2 = i, maxIdx // 记录局部交换的下标
		} else if s[i] > s[maxIdx] {
			maxIdx = i
		}
	}

	if idx1 < 0 {
		return num
	}

	// 交换最终结果
	s[idx1], s[idx2] = s[idx2], s[idx1]

	result, _ := strconv.Atoi(string(s))
	return result
}

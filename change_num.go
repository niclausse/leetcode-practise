package leetcode_practise

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

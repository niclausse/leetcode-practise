package sort

// BubbleSort 冒泡排序, 从第一位开始, 不断与后一位比较大小, 每次将最大元素放到最后
// 时间复杂度: O(n^2), n为数组长度
// 最好情况下, arr已经是有序排列, 则只需要进行一次冒泡比较, 时间复杂度为O(n)
// 最坏情况下, arr元素为倒序排列, 则需要进行n次冒泡比较, 时间复杂度为O(n^2)
func BubbleSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	for i := 0; i < len(arr); i++ {
		var swapped bool
		for j := 0; j < len(arr)-i-1; j++ { // 相邻位两两比较
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true // 有数据交换
			}
		}

		// 如果没有数据交换, 说明0～len(arr)-i-1内元素已经有序, 不再需要继续冒泡
		if !swapped {
			break
		}
	}
}

package sort

// QuickSort 快速排序: 选出个基准值pivot, 小于基准值的元素放到左边, 大于基准值的放到右边
func QuickSort(arr []int, start, end int) {
	if start < end {
		pi := quickPartitionV2(arr, start, end)
		QuickSort(arr, start, pi-1)
		QuickSort(arr, pi+1, end)
	}
}

func quickPartitionV2(arr []int, startIndex, endIndex int) int {
	var (
		pivot = arr[startIndex]
		left  = startIndex
		right = endIndex
	)

	for left != right {
		for left < right && pivot < arr[right] {
			right--
		}

		for left < right && pivot >= arr[left] {
			left++
		}

		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}

	arr[startIndex], arr[left] = arr[left], arr[startIndex]

	return left
}

func quickPartition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[high] = arr[high], arr[i]
	return i
}

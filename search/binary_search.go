package search

// BinarySearch 二分搜索
// arr中存在value时, 返回对应下标; 否则返回-1
func BinarySearch(arr []int, value int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := low + (high-low)/2

		if arr[mid] == value {
			return mid
		}

		if arr[mid] > value {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

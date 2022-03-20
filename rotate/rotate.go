package rotate

// 题49
// 矩阵翻转90度
// 水平翻转 + 对角线翻转
// 5  1  9  11
// 2  4  8  10
// 13 3  6  7
// 15 14 12 16
func rotate(matrix [][]int) {
	// 水平翻转
	for i := 0; i < len(matrix)>>1; i++ {
		matrix[i], matrix[len(matrix)-1-i] = matrix[len(matrix)-1-i], matrix[i]
	}

	// 对角线翻转
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

// 使用辅助数组
func rotate1(matrix [][]int) {
	l2 := make([][]int, 0, len(matrix))

	for j := 0; j < len(matrix); j++ {
		l1 := make([]int, 0, len(matrix))
		for i := len(matrix) - 1; i >= 0; i-- {
			l1 = append(l1, matrix[i][j])
		}

		l2 = append(l2, l1)
	}

	copy(matrix, l2)
}

package leetcode_practise

import "strings"

// 1～3999
func convert(num int) string {
	hashTable := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}

	arr := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	var results []string
	count := 0

	for i := 0; i < len(arr) && num != 0; i++ {
		count = num / arr[i]
		num = num - count*arr[i]
		for count != 0 {
			results = append(results, hashTable[arr[i]])
			count--
		}
	}
	return strings.Join(results, "")
}

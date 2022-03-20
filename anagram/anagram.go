package anagram

import "sort"

// 简单单词串场景：s, t 只考虑是小写的英文单词
func simpleAnagram(s, t string) bool {
	// 异位词等价于两个字符串中字符出现的字符的种类和数量均相等。
	// 由于字符串只包含26个小写英文字母，因此可以维护两个长度为26的频次数组，分别记录两个字符串中字符的频次
	var c1, c2 [26]int
	for _, c := range s {
		c1[c-'a']++
	}

	for _, c := range t {
		c2[c-'a']++
	}

	return c1 == c2
}

// 使用hashTable代替simple场景中的数组记录字符出现的频次即可
func unicodeAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	hashTable := make(map[rune]int)
	for _, c := range s {
		hashTable[c]++
	}

	for _, c := range t {
		hashTable[c]--
		if hashTable[c] < 0 {
			return false
		}
	}
	return true
}

// 异位词分组
// 思路1：使用字符串中字符出现的频次作为hash_key, 相同频次的字符串数组作为hash_value
// 思路2：使用排序后的字符串作为hash_key
func anagramGroup(arr []string) [][]string {
	hashTable := make(map[[26]int][]string)

	for _, s := range arr {
		key := [26]int{}
		for _, c := range s {
			key[c-'a']++
		}
		hashTable[key] = append(hashTable[key], s)
	}

	groups := make([][]string, 0, len(hashTable))
	for _, group := range hashTable {
		groups = append(groups, group)
	}

	return groups
}

func anagramGroupWithSort(arr []string) [][]string {
	hashTable := make(map[string][]string)
	for _, s := range arr {
		tmp := []byte(s)
		sort.Slice(tmp, func(i, j int) bool {
			return tmp[i] < tmp[j]
		})

		sorted := string(tmp)
		hashTable[sorted] = append(hashTable[sorted], s)
	}

	groups := make([][]string, 0, len(hashTable))
	for _, g := range hashTable {
		groups = append(groups, g)
	}

	return groups
}

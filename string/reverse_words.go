package string

import (
	"strings"
)

func reverseWord(s string) (res string) {
	slice := strings.Split(strings.TrimSpace(s), " ")

	for i := 0; i < len(slice)>>1; i++ {
		slice[i], slice[len(slice)-1-i] = slice[len(slice)-1-i], slice[i]
	}

	var words []string
	for _, tmp := range slice {
		if tmp != "" {
			words = append(words, tmp)
		}
	}

	return strings.Join(words, " ")
}

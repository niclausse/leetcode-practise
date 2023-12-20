package string

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	fmt.Println(reverseWord("It's a     word!   "))
}

func TestValid(t *testing.T) {
	s := ",1ab"
	for _, a := range s {
		t.Log(valid(a))
	}
}

func TestIsPalindrome(t *testing.T) {
	t.Log(isPalindrome(",12!21"))
}

func TestLcp(t *testing.T) {
	pre := lcp("flows", "flower")
	if pre != "flow" {
		t.Errorf("longest conmmon prefix should be 'flow'")
		return
	}
}

func TestLongestPalindrome(t *testing.T) {
	s := "cbbd"
	p := longestPalindrome(s)
	if p != "bb" {
		t.Errorf("cbbd's palindrome should be bb")
	}
	fmt.Println(p)
}

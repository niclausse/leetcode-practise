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

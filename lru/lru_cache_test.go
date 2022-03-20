package lru

import (
	"fmt"
	"testing"
)

func TestCache(t *testing.T) {
	cache := Constructor(2)
	fmt.Println(cache.Get(2))

	cache.Put(2, 6)
	fmt.Println(cache.Get(1))

	cache.Put(1, 5)
	cache.Put(1, 2)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(2))
}

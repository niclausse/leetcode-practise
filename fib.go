package leetcode_practise

// 斐波拉契数
// 0 1 1 2 3 5 ...
func fib(n int) int {
	if n < 2 {
		return n
	}

	p, q, r := 0, 0, 1

	for i := 2; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}

	return r
}
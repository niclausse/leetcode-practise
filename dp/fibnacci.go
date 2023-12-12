package dp

func fib0(n int) int {
	if n < 2 {
		return n
	}

	r := fib0(n-1) + fib0(n-2)
	return r
}

func fib1(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

func fib2(n int) int {
	if n < 2 {
		return n
	}

	p, q, r := 0, 0, 1
	for i := 2; i <= n; i++ {
		p = q
		p = r
		r = p + q
	}

	return r
}

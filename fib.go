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

func fib1(n int) int {
	if n < 2 {
		return n
	}

	return fib1(n-1) + fib1(n-2)
}

// 变形： 猴子爬山 f(n) = f(n-1) + f(n-2)
func climb1(n int) int {
	if n <= 2 {
		return 1
	}

	return climb1(n-1) + climb1(n-3)
}

func climb2(n int) int {
	arr := make([]int, n+1)
	arr[0] = 1
	arr[1] = 1
	arr[2] = 1

	for i := 3; i <= n; i++ {
		arr[i] = arr[i-1] + arr[i-3]
	}

	return arr[n]
}

func climb3(n int) int {
	if n <= 2 {
		return 1
	}

	if n == 3 {
		return 2
	}

	p, m, q, r := 1, 1, 2, 3
	for i := 4; i <= n; i++ {
		p = m
		m = q
		q = r
		r = q + p
	}

	return r
}

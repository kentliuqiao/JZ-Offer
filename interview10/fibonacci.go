package interview10

func fibonacciRecursive(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return fibonacciRecursive(n-1) + fibonacciRecursive(n-2)
}

func fibonacciIterative(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}

	return x
}

// Note: Fibonacci数列还有一种O(longn)时间复杂度的解法。需要用到数学公式以及一个递归求数值次方的算法

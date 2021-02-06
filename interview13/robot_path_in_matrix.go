package interview13

func movingCount(m, n int, k int) int {
	if k < 0 || m <= 0 || n <= 0 {
		return 0
	}
	visited := make([]bool, m*n)
	for i := range visited {
		visited[i] = false
	}

	count := dfs(m, n, k, 0, 0, visited)

	return count
}

func dfs(k int, rows, cols int, row, col int, visited []bool) int {
	count := 0
	if check(k, rows, cols, row, col, visited) {
		visited[row*cols+col] = true
		count = 1 +
			dfs(k, rows, cols, row+1, col, visited) +
			dfs(k, rows, cols, row, col+1, visited) +
			dfs(k, rows, cols, row, col-1, visited) +
			dfs(k, rows, cols, row-1, col, visited)
	}

	return count
}

func check(k int, rows, cols int, row, col int, visited []bool) bool {
	if row >= 0 && row < rows && col >= 0 && col < cols && getDigitSum(row)+getDigitSum(col) <= k && !visited[row*cols+col] {
		return true
	}

	return false
}

func getDigitSum(number int) int {
	sum := 0
	for number > 0 {
		sum += number % 10
		number /= 10
	}

	return sum
}

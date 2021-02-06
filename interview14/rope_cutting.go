package interview14

import "math"

func maxProductDP(n int) int {
	if n < 2 {
		return 0
	}
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}

	products := make([]int, n)
	products[0] = 0
	products[1] = 0
	products[2] = 1
	products[3] = 2

	for i := 4; i <= n; i++ {
		max := 0
		for j := 1; j <= i/2; j++ {
			product := products[j] * products[i-j]
			if max < product {
				max = product
			}
		}
		products[i] = max
	}

	return products[n]
}

func maxProductGreedy(n int) int {
	if n < 2 {
		return 0
	}
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}

	timesOf3 := n / 3
	if n-timesOf3*3 == 1 {
		timesOf3--
	}

	timesOf2 := (n - timesOf3*3) / 2
	return int(math.Pow(3, float64(timesOf3))) * int(math.Pow(2, float64(timesOf2)))
}

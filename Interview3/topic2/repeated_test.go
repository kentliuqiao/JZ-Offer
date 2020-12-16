package topic2

import "testing"

var cases = []struct {
	nums []int
	want []int
}{
	{
		nums: []int{2, 3, 5, 4, 3, 2, 6, 7},
		want: []int{2, 3},
	},
	{
		nums: []int{},
		want: []int{-1},
	},
	{
		nums: []int{1, 3, 5, 4, 3, 2, 6, 7},
		want: []int{3},
	},
	// {
	// 	nums: []int{2, 3, 5, 4, 3, 2, 6, 7},
	// 	want: 2,
	// },
}

func TestSolution1(t *testing.T) {
	for _, c := range cases {
		got := solution1(c.nums)
		if !inSlice(got, c.want) {
			t.Errorf("got: %d, want: %v", got, c.want)
		}
	}
}

func TestSolution2(t *testing.T) {
	for _, c := range cases {
		got := solution1(c.nums)
		if !inSlice(got, c.want) {
			t.Errorf("got: %d, want: %v", got, c.want)
		}
	}
}

func inSlice(target int, arr []int) bool {
	for _, elm := range arr {
		if elm == target {
			return true
		}
	}

	return false
}

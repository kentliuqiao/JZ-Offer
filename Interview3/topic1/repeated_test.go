package topic1

import "testing"

var cases = []struct {
	nums []int
	num  int
}{
	{
		nums: []int{2, 3, 1, 0, 2, 5, 3},
		num:  2,
	},
	{
		nums: []int{4, 3, 1, 0, 2, 5, 3},
		num:  3,
	},
	{
		nums: []int{0, 3, 1, 0, 2, 5, -1},
		num:  -1,
	},
	{
		nums: []int{0, 3, 1, 0, 2, 5, 7},
		num:  -1,
	},
	{
		nums: []int{},
		num:  -1,
	},
}

func TestSolution3(t *testing.T) {
	for _, c := range cases {
		got := Solution3(c.nums)
		if c.num != got {
			t.Errorf("arr: %v, want: %d, got: %d", c.nums, c.num, got)
		}
	}
}

func TestSolution4(t *testing.T) {
	for _, c := range cases {
		got := Solution3(c.nums)
		if c.num != got {
			t.Errorf("arr: %v, want: %d, got: %d", c.nums, c.num, got)
		}
	}
}

package interview4

import (
	"testing"
)

var cases = []struct {
	matrix [][]int
	target int
	want   bool
}{
	{
		matrix: [][]int{
			{1, 2, 8, 9},
			{2, 4, 9, 12},
			{4, 7, 10, 13},
			{6, 8, 11, 15},
		},
		target: 7,
		want:   true,
	},
	{
		matrix: [][]int{
			{1, 2, 8, 9},
			{2, 4, 9, 12},
			{4, 7, 10, 13},
			{6, 8, 11, 15},
		},
		target: 5,
		want:   false,
	},
}

func TestSearchMatrix(t *testing.T) {
	for _, c := range cases {
		got := searchMatrix(c.matrix, c.target)
		if c.want != got {
			t.Errorf("matrix: %+v; want: %v; got: %v", c.matrix, c.want, got)
		}
	}
}

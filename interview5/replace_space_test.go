package interview5

import "testing"

func TestReplaceSpace(t *testing.T) {
	cases := []struct {
		str  string
		want string
	}{
		{
			str:  "we are happy.",
			want: "we%20are%20happy.",
		},
		{
			str:  "",
			want: "",
		},
		{
			str:  "u  are happy.",
			want: "u%20%20are%20happy.",
		},
	}

	for _, c := range cases {
		got := replaceSpace(c.str)
		if got != c.want {
			t.Errorf("TestReplaceSpace failed, want: %s, got :%s", c.want, got)
		}
	}
}

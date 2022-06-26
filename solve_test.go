package rpn

import "testing"

func TestSolve(t *testing.T) {
	tt := []struct {
		got  string
		want int
	}{
		{"3 + 4", 7},
		{"3 + 4 * 2 / ( 1 - 5 ) ^ 2 ^ 3", 3},
	}

	for _, tc := range tt {
		oq, err := Parse(tc.got)
		if err != nil {
			t.Error(err)
		}
		num, err := Solve(oq)
		if err != nil {
			t.Error(err)
		}
		if num != tc.want {
			t.Errorf("got = %d, want = %d", num, tc.want)
		}
	}
}

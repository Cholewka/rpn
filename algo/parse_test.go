// Algo is an implementation of the shunting yard algorithm.
// https://en.wikipedia.org/wiki/Shunting_yard_algorithm
package algo

import "testing"

func TestParse(t *testing.T) {
	got, err := Parse("3 + 4 * 2 / ( 1 - 5 ) ^ 2 ^ 3")
	if err != nil {
		t.Errorf("Parse() returned an error: %v\n", err)
	}
	want := []string{"3", "4", "2", "*", "1", "5", "-", "2", "3", "^", "^"}
	for i, v := range want {
		if got[i] != v {
			t.Errorf("want[%d] = %q, got[%d] = %q", i, v, i, got[i])
		}
	}
}

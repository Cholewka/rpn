package rpn

import "testing"

func TestParse(t *testing.T) {
	got, err := Parse("3 + 4 * 2 / ( 1 - 5 ) ^ 2 ^ 3")
	if err != nil {
		t.Errorf("Parse() returned an error: %v\n", err)
	}
	want := []string{"3", "4", "2", "*", "1", "5", "-", "2", "3", "^", "^", "/", "+"}
	if len(want) > len(got) {
		t.Errorf("len mismatch between got and want")
	}
	for i, v := range got {
		if want[i] != v {
			t.Errorf("want[%d] = %q, got[%d] = %q", i, got[i], i, want)
		}
	}
}

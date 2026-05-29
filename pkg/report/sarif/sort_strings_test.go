//ff:func feature=report type=test control=iteration dimension=1 topic=sarif
//ff:what TestSortStrings — in-place 오름차순 정렬 (정상/이미정렬/빈슬라이스) 검증
package sarif

import (
	"testing"
)

// TestSortStrings checks the insertion sort orders in place and handles
// already-sorted, reversed, duplicate, and empty inputs.
func TestSortStrings(t *testing.T) {
	cases := []struct {
		name string
		in   []string
		want []string
	}{
		{"reversed", []string{"c", "b", "a"}, []string{"a", "b", "c"}},
		{"already sorted", []string{"a", "b", "c"}, []string{"a", "b", "c"}},
		{"duplicates", []string{"b", "a", "b"}, []string{"a", "b", "b"}},
		{"single", []string{"x"}, []string{"x"}},
		{"empty", []string{}, []string{}},
		{"rule ids", []string{"XOS-15", "S-27", "M-2"}, []string{"M-2", "S-27", "XOS-15"}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// Copy in place without changing nil-ness vs empty-ness.
			s := make([]string, len(c.in))
			copy(s, c.in)
			sortStrings(s)
			if len(s) != len(c.want) {
				t.Fatalf("len: got %d, want %d", len(s), len(c.want))
			}
			for i := range c.want {
				if s[i] != c.want[i] {
					t.Errorf("index %d: got %q, want %q", i, s[i], c.want[i])
				}
			}
		})
	}
}

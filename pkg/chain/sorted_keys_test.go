//ff:func feature=chain type=test control=iteration dimension=1
//ff:what TestSortedKeys — sortedStringKeys 가 map 키를 정렬 반환하는지 검증 (empty/single/multi)

package chain

import (
	"reflect"
	"testing"
)

func TestSortedKeys(t *testing.T) {
	cases := []struct {
		name string
		in   map[string]bool
		want []string
	}{
		{name: "empty", in: map[string]bool{}, want: []string{}},
		{name: "single", in: map[string]bool{"only": true}, want: []string{"only"}},
		{
			name: "multi unsorted",
			in:   map[string]bool{"delta": true, "alpha": true, "charlie": true, "bravo": true},
			want: []string{"alpha", "bravo", "charlie", "delta"},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := sortedStringKeys(tc.in)
			if len(got) != len(tc.want) {
				t.Fatalf("len: got %d (%v), want %d (%v)", len(got), got, len(tc.want), tc.want)
			}
			if len(tc.want) > 0 && !reflect.DeepEqual(got, tc.want) {
				t.Errorf("sortedStringKeys = %v, want %v", got, tc.want)
			}
		})
	}
}

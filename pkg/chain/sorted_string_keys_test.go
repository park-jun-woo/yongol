//ff:func feature=chain type=test control=sequence
//ff:what sortedStringKeys 가 map 키를 정렬해 반환하는지 검증
package chain

import (
	"reflect"
	"testing"
)

func TestSortedStringKeys(t *testing.T) {
	if got := sortedStringKeys(map[string]bool{}); len(got) != 0 {
		t.Errorf("empty map: got %v, want empty", got)
	}

	in := map[string]bool{"charlie": true, "alpha": true, "bravo": true}
	want := []string{"alpha", "bravo", "charlie"}
	if got := sortedStringKeys(in); !reflect.DeepEqual(got, want) {
		t.Errorf("sortedStringKeys = %v, want %v", got, want)
	}
}

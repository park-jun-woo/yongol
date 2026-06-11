//ff:func feature=gen-react type=test control=sequence
//ff:what dedupeStrings — 첫 등장 순서 보존 중복 제거/빈 입력 검증

package react

import (
	"reflect"
	"testing"
)

func TestDedupeStrings(t *testing.T) {
	got := dedupeStrings([]string{"/b/", "/a", "/b/", "/c", "/a"})
	want := []string{"/b/", "/a", "/c"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("dedupeStrings = %v, want %v", got, want)
	}

	if got := dedupeStrings(nil); got != nil {
		t.Errorf("nil in, nil out — got %v", got)
	}
}

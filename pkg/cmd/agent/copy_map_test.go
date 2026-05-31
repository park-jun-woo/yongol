//ff:func feature=agent type=test control=sequence
//ff:what TestCopyMap — copyMap 이 동일 내용의 독립 복사본을 만드는지 검증
package agent

import (
	"reflect"
	"testing"
)

func TestCopyMap(t *testing.T) {
	src := map[string]any{"a": 1, "b": "x"}
	got := copyMap(src)
	if !reflect.DeepEqual(got, src) {
		t.Fatalf("copy mismatch: got %v want %v", got, src)
	}
	// Mutating the copy must not affect the source.
	got["a"] = 99
	if src["a"] != 1 {
		t.Errorf("mutation leaked to source: %v", src["a"])
	}

	empty := copyMap(map[string]any{})
	if len(empty) != 0 {
		t.Errorf("empty copy should be empty, got %v", empty)
	}
}

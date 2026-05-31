//ff:func feature=gen-splitter type=test control=sequence
//ff:what snake / needsSnakeBreak / isVersionSuffix / tailSegment / summariseDoc 순수 헬퍼
package splitter

import (
	"testing"
)

func TestTailSegment(t *testing.T) {
	if got := tailSegment("/a/b/c.go"); got != "c.go" {
		t.Errorf("tailSegment = %q, want c.go", got)
	}
	if got := tailSegment("plain.go"); got != "plain.go" {
		t.Errorf("tailSegment = %q, want plain.go", got)
	}
}

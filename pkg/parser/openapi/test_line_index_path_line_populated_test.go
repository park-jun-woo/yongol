//ff:func feature=openapi-parse type=test control=sequence
//ff:what LineIndex — 실 fixture 기반 PathLine lookup 검증

package openapi

import "testing"

func TestLineIndex_PathLine_Populated(t *testing.T) {
	path := writeFixture(t)
	idx, err := BuildLineIndex(path)
	if err != nil {
		t.Fatalf("BuildLineIndex: %v", err)
	}
	if got := idx.PathLine("/login"); got != 14 {
		t.Errorf("PathLine(/login) = %d, want 14", got)
	}
	if got := idx.PathLine("/missing"); got != 0 {
		t.Errorf("PathLine(/missing) = %d, want 0", got)
	}
}

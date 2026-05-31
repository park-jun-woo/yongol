//ff:func feature=agent type=test control=sequence
//ff:what TestMatchByPath — 에러 메시지에 포함된 path 키로 op 매핑 검증
package agent

import (
	"testing"
)

func TestMatchByPath(t *testing.T) {
	offsets := []pathOffset{
		{Op: "OpA", Path: "/users/{id}"},
		{Op: "OpB", Path: "/orders"},
		{Op: "OpC", Path: ""}, // empty path must never match
	}
	got := matchByPath("error at /orders endpoint", offsets)
	if len(got) != 1 || got[0] != "OpB" {
		t.Errorf("matchByPath = %v, want [OpB]", got)
	}
	if got := matchByPath("no path here", offsets); len(got) != 0 {
		t.Errorf("expected no match, got %v", got)
	}
}

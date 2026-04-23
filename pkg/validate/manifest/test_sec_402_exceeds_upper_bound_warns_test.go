//ff:func feature=validate type=test control=sequence topic=manifest-auth
//ff:what SEC-402 테스트 — 2h 는 상한 초과 WARNING

package manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestSEC402_ExceedsUpperBound_Warns(t *testing.T) {
	diags := sec402AccessTTLUpperBound(fsWithAccessTTL("2h"))
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d: %+v", len(diags), diags)
	}
	if diags[0].Level != diagnostic.LevelWarning {
		t.Fatalf("expected WARNING, got %v", diags[0].Level)
	}
}

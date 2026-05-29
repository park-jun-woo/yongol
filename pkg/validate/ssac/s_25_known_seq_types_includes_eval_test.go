//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what S-25 정합성 — knownSeqTypes 가 @eval 을 포함 (BUG-001 명시 가드)

package ssac

import (
	"testing"

	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// TestS25KnownSeqTypesIncludesEval is an explicit guard for BUG-001. Even if
// the parser export shape changes, @eval must remain whitelisted.
func TestS25KnownSeqTypesIncludesEval(t *testing.T) {
	if !knownSeqTypes[parsessac.SeqEval] {
		t.Fatalf("knownSeqTypes must contain %q (BUG-001 regression)", parsessac.SeqEval)
	}
}

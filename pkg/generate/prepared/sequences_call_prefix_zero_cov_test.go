//ff:func feature=generate type=test control=sequence
//ff:what prepared 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package prepared

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestSequencesCallPrefix_ZeroCov(t *testing.T) {
	seqs := []ssac.Sequence{{Type: "call", Model: "session.Put"}, {Type: "get"}}
	if !sequencesCallPrefix(seqs, "session.") {
		t.Error("expected prefix match")
	}
	if sequencesCallPrefix(seqs, "cache.") {
		t.Error("unexpected match")
	}
}

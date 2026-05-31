//ff:func feature=migration type=test control=sequence
//ff:what 마이그레이션 파이프라인 함수별 named 테스트 — tsma 함수명 매칭용 (parse/diff/emit/tokenizer 커버)
package migration

import (
	"testing"
)

func TestRenameMaps(t *testing.T) {
	h := newEmptyHints()
	h.RenameTables = []RenameTableHint{{From: "a", To: "b"}}
	fwd, rev := renameMaps(h)
	if fwd["a"] != "b" || rev["b"] != "a" {
		t.Errorf("rename maps wrong: %v %v", fwd, rev)
	}
	// nil hints → empty maps.
	if f, r := renameMaps(nil); len(f) != 0 || len(r) != 0 {
		t.Errorf("nil hints should give empty maps")
	}
}

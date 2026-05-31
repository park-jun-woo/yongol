//ff:func feature=migration type=test control=sequence
//ff:what 마이그레이션 파이프라인 함수별 named 테스트 — tsma 함수명 매칭용 (parse/diff/emit/tokenizer 커버)
package migration

import (
	"testing"
)

func TestCollectRenameOps(t *testing.T) {
	h := newEmptyHints()
	h.RenameTables = []RenameTableHint{{From: "a", To: "b"}}
	h.RenameColumns = []RenameColumnHint{{Table: "b", From: "x", To: "y"}}
	if len(collectRenameOps(h)) == 0 {
		t.Errorf("expected rename ops")
	}
}

//ff:func feature=migration type=test control=sequence
//ff:what 마이그레이션 파이프라인 함수별 named 테스트 — tsma 함수명 매칭용 (parse/diff/emit/tokenizer 커버)
package migration

import (
	"testing"
)

func TestLookupPrevColumn(t *testing.T) {
	col := &Column{Name: "email"}
	prevMap := map[string]*Column{"email": col}
	if got := lookupPrevColumn("email", prevMap, map[string]bool{}, newEmptyHints(), "users"); got != col {
		t.Errorf("expected direct lookup hit")
	}
	if got := lookupPrevColumn("missing", prevMap, map[string]bool{}, newEmptyHints(), "users"); got != nil {
		t.Errorf("missing column should be nil")
	}
}

//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestFindCallInStmt — 매칭 호출 발견 / 비매칭 nil / 첫 매칭 단락 검증
package qcheck

import (
	"testing"
)

func TestFindCallInStmt_NotFound(t *testing.T) {
	list := blockStmts(t, "_ = json.Marshal(v)")
	if got := findCallInStmt(list[0], "json", "Unmarshal"); got != nil {
		t.Errorf("expected nil for non-matching call, got %+v", got)
	}
}

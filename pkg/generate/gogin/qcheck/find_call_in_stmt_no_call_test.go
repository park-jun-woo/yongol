//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestFindCallInStmt — 매칭 호출 발견 / 비매칭 nil / 첫 매칭 단락 검증
package qcheck

import (
	"testing"
)

func TestFindCallInStmt_NoCall(t *testing.T) {
	list := blockStmts(t, "x := 1")
	if got := findCallInStmt(list[0], "json", "Unmarshal"); got != nil {
		t.Errorf("expected nil when no call present, got %+v", got)
	}
}

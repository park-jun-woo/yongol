//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestCollectResourceFromBlock — top-level + 중첩 if-블록 내 미닫힘 리소스 DF-06 재귀 수집
package qcheck

import (
	"testing"
)

func TestCollectResourceFromBlock_TopLevelClosed(t *testing.T) {
	// Resource acquired at top level WITH defer Close -> no finding.
	src := `package x
func H() {
	f, err := os.Open("x")
	defer f.Close()
	_ = err
}`
	body, fset := parseFuncBody(t, src)
	if got := collectResourceFromBlock(body, fset); len(got) != 0 {
		t.Errorf("expected no findings when closed, got %+v", got)
	}
}

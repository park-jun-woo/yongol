//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestCollectResourceFromBlock — top-level + 중첩 if-블록 내 미닫힘 리소스 DF-06 재귀 수집
package qcheck

import (
	"testing"
)

func TestCollectResourceFromBlock_NestedMissing(t *testing.T) {
	// os.Open inside an if-block with no defer Close -> one DF-06 from recursion.
	src := `package x
func H(cond bool) {
	if cond {
		f, err := os.Open("x")
		_ = err
		_ = f
	}
}`
	body, fset := parseFuncBody(t, src)
	findings := collectResourceFromBlock(body, fset)
	if len(findings) != 1 {
		t.Fatalf("want 1 DF-06 finding from nested block, got %d: %+v", len(findings), findings)
	}
	if findings[0].Category != "DF-06" || findings[0].Detail != "os.Open" {
		t.Errorf("unexpected finding: %+v", findings[0])
	}
}

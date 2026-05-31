//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what TestCollectScanFromBlock — 중첩 블록 내 미가드 .Scan() DF-02 재귀 수집 + 가드 케이스
package qcheck

import (
	"testing"
)

func TestCollectScanFromBlock_NestedUnguarded(t *testing.T) {
	src := `package x
func H(cond bool, r row) {
	if cond {
		_ = r.Scan(nil)
	}
}`
	body, fset := parseHBody(t, src)
	findings := collectScanFromBlock(body, fset)
	if len(findings) == 0 {
		t.Fatalf("want at least 1 DF-02 finding from nested block, got none")
	}
	for _, f := range findings {
		if f.Category != "DF-02" {
			t.Errorf("unexpected finding category: %+v", f)
		}
	}
}

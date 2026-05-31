//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestCollectScanFromBlock — 중첩 블록 내 미가드 .Scan() DF-02 재귀 수집 + 가드 케이스
package qcheck

import (
	"testing"
)

func TestCollectScanFromBlock_Guarded(t *testing.T) {
	src := `package x
func H(r row) {
	err := r.Scan(nil)
	if err != nil { return }
}`
	body, fset := parseHBody(t, src)
	if got := collectScanFromBlock(body, fset); len(got) != 0 {
		t.Errorf("expected no findings for guarded Scan, got %+v", got)
	}
}

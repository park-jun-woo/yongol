//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestWarnExceeds_Clean — 위반 없는 소스에 대해 WarnExceeds가 빈 슬라이스 반환

package qcheck

import "testing"

func TestWarnExceeds_Clean(t *testing.T) {
	warns := WarnExceeds("clean.go", cleanSrc, DefaultLimits())
	if len(warns) != 0 {
		t.Fatalf("want 0 warns on clean src, got %v", warns)
	}
}

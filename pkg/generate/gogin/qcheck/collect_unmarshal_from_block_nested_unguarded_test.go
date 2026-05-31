//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what TestCollectUnmarshalFromBlock — 중첩 블록 내 미가드 json.Unmarshal DF-01 재귀 + 가드 케이스
package qcheck

import (
	"testing"
)

func TestCollectUnmarshalFromBlock_NestedUnguarded(t *testing.T) {
	src := `package x
func H(cond bool, b []byte, v any) {
	if cond {
		_ = json.Unmarshal(b, v)
	}
}`
	body, fset := unmarshalBody(t, src)
	findings := collectUnmarshalFromBlock(body, []string{"json"}, fset)
	if len(findings) == 0 {
		t.Fatalf("want at least 1 DF-01 finding from nested block, got none")
	}
	for _, f := range findings {
		if f.Category != "DF-01" || f.Detail != "json.Unmarshal" {
			t.Errorf("unexpected finding: %+v", f)
		}
	}
}

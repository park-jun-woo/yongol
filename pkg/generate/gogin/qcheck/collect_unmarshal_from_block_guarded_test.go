//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestCollectUnmarshalFromBlock — 중첩 블록 내 미가드 json.Unmarshal DF-01 재귀 + 가드 케이스
package qcheck

import (
	"testing"
)

func TestCollectUnmarshalFromBlock_Guarded(t *testing.T) {
	src := `package x
func H(b []byte, v any) {
	if err := json.Unmarshal(b, v); err != nil { return }
}`
	body, fset := unmarshalBody(t, src)
	if got := collectUnmarshalFromBlock(body, []string{"json"}, fset); len(got) != 0 {
		t.Errorf("expected no findings for guarded Unmarshal, got %+v", got)
	}
}

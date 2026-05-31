//ff:func feature=migration type=test control=iteration dimension=1
//ff:what 마이그레이션 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package migration

import (
	"testing"
)

func TestStepTopLevel_ZeroCov(t *testing.T) {
	st := newSplitState()
	s := "a,b"
	i := 0
	for i < len(s) {
		i = stepTopLevel(st, s, i, ',')
		i++
	}
	out := st.finish()
	if len(out) != 2 {
		t.Errorf("expected 2 parts, got %v", out)
	}
}

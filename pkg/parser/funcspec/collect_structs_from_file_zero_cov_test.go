//ff:func feature=funcspec type=test control=sequence
//ff:what funcspec 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package funcspec

import (
	"testing"
)

func TestCollectStructsFromFile_ZeroCov(t *testing.T) {
	_, f := bnFile(t, bnStructSrc)
	result := map[string][]Field{}
	collectStructsFromFile(f, result)
	if len(result) != 2 {
		t.Errorf("expected 2 structs, got %d", len(result))
	}
}

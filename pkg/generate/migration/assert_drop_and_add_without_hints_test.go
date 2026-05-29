//ff:func feature=migration type=test-helper control=iteration dimension=1
//ff:what assertDropAndAddWithoutHints — 힌트 없을 때 Drop+Add 쌍이 생성됐는지 검증
package migration

import "testing"

func assertDropAndAddWithoutHints(t *testing.T, prev, curr *Schema) {
	t.Helper()
	opsNoHint := Diff(prev, curr, nil)
	hasDrop, hasAdd := false, false
	for _, op := range opsNoHint {
		if _, ok := op.(DropColumn); ok {
			hasDrop = true
		}
		if _, ok := op.(AddColumn); ok {
			hasAdd = true
		}
	}
	if !hasDrop || !hasAdd {
		t.Errorf("without hints expected Drop + Add; got %+v", opsNoHint)
	}
}

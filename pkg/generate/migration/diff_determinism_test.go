//ff:func feature=migration type=test control=iteration dimension=1
//ff:what TestDiff_Determinism — 같은 입력 2회 호출 결과가 같은 순서로 산출되는지 확인
package migration

import "testing"

func TestDiff_Determinism(t *testing.T) {
	prev := mustAST(t, `CREATE TABLE t (id INTEGER);`)
	curr := mustAST(t, `CREATE TABLE t (id INTEGER, name TEXT, age INTEGER);`)
	a := Diff(prev, curr, nil)
	b := Diff(prev, curr, nil)
	if len(a) != len(b) {
		t.Fatalf("different lengths: %d vs %d", len(a), len(b))
	}
	for i := range a {
		if a[i].Description() != b[i].Description() {
			t.Errorf("order differs at %d: %q vs %q", i, a[i].Description(), b[i].Description())
		}
	}
}

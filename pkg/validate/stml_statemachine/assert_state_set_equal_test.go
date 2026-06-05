//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=stml-statemachine
//ff:what assertStateSetEqual — 단일 Symbol의 상태집합(got/want)이 동일한지 비교

package stml_statemachine

import "testing"

// assertStateSetEqual fails the test unless gotSet contains every state in
// wantSet for the given symbol. Callers compare lengths before calling.
func assertStateSetEqual(t *testing.T, sym string, gotSet, wantSet map[string]bool) {
	t.Helper()
	for state := range wantSet {
		if !gotSet[state] {
			t.Errorf("symbol %q missing state %q (got %+v)", sym, state, gotSet)
		}
	}
}

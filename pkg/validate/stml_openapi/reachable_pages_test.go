//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-openapi
//ff:what reachablePages — 루트에서 간선 전파 / 미도달 격리 / 고스트 루트 무시 BFS 검증

package stml_openapi

import "testing"

func TestReachablePages(t *testing.T) {
	g := &pageGraph{
		Pages: []string{"a", "b", "c", "d"},
		Roots: map[string]bool{"a": true, "ghost": true},
		Edges: map[string][]string{
			"a": {"b"},
			"b": {"c"},
			"d": {"c"}, // edge from an unreachable source must not seed anything
		},
	}
	reached := reachablePages(g)
	for _, name := range []string{"a", "b", "c"} {
		if !reached[name] {
			t.Errorf("%q should be reachable", name)
		}
	}
	if reached["d"] {
		t.Error("d has no incoming path from the roots and must stay unreachable")
	}
	if reached["ghost"] {
		t.Error("a root naming no page must not enter the reached set")
	}
}

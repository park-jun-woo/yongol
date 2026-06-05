//ff:func feature=validate type=test control=iteration dimension=2 topic=stml-openapi
//ff:what TestCollectEachComponentNames — 직접 ComponentRefs와 children 재귀 두 경로 수집 검증

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestCollectEachComponentNames(t *testing.T) {
	e := stml.EachBlock{
		Field:      "items",
		Components: []stml.ComponentRef{{Name: "DirectComp"}},
		Children: []stml.ChildNode{
			{Kind: "component", Component: &stml.ComponentRef{Name: "ChildComp"}},
		},
	}

	out := map[string]struct{}{}
	collectEachComponentNames(e, out)

	want := []string{"DirectComp", "ChildComp"}
	for _, name := range want {
		if _, ok := out[name]; !ok {
			t.Errorf("missing component name %q", name)
		}
	}
	if len(out) != len(want) {
		t.Errorf("expected %d names, got %d: %+v", len(want), len(out), out)
	}
}

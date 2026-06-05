//ff:func feature=validate type=test control=iteration dimension=3 topic=stml-openapi
//ff:what TestCollectPageComponentNames — Fetches/Actions/Children 세 경로 수집 검증

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestCollectPageComponentNames(t *testing.T) {
	page := stml.PageSpec{
		Name: "page",
		Fetches: []stml.FetchBlock{{
			Components: []stml.ComponentRef{{Name: "FetchComp"}},
		}},
		Actions: []stml.ActionBlock{{
			Fields: []stml.FieldBind{{Tag: "data-component:ActionComp"}},
		}},
		Children: []stml.ChildNode{
			{Kind: "component", Component: &stml.ComponentRef{Name: "ChildComp"}},
		},
	}

	out := map[string]struct{}{}
	collectPageComponentNames(page, out)

	want := []string{"FetchComp", "ActionComp", "ChildComp"}
	for _, name := range want {
		if _, ok := out[name]; !ok {
			t.Errorf("missing component name %q", name)
		}
	}
	if len(out) != len(want) {
		t.Errorf("expected %d names, got %d: %+v", len(want), len(out), out)
	}
}

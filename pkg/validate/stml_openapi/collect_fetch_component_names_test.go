//ff:func feature=validate type=test control=iteration dimension=4 topic=stml-openapi
//ff:what TestCollectFetchComponentNames — Components/Eaches/NestedFetches/Children 네 경로 수집 검증

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestCollectFetchComponentNames(t *testing.T) {
	f := stml.FetchBlock{
		OperationID: "ListItems",
		Components:  []stml.ComponentRef{{Name: "DirectComp"}},
		Eaches: []stml.EachBlock{{
			Components: []stml.ComponentRef{{Name: "EachComp"}},
		}},
		NestedFetches: []stml.FetchBlock{{
			Components: []stml.ComponentRef{{Name: "NestedComp"}},
		}},
		Children: []stml.ChildNode{
			{Kind: "component", Component: &stml.ComponentRef{Name: "ChildComp"}},
		},
	}

	out := map[string]struct{}{}
	collectFetchComponentNames(f, out)

	want := []string{"DirectComp", "EachComp", "NestedComp", "ChildComp"}
	for _, name := range want {
		if _, ok := out[name]; !ok {
			t.Errorf("missing component name %q", name)
		}
	}
	if len(out) != len(want) {
		t.Errorf("expected %d names, got %d: %+v", len(want), len(out), out)
	}
}

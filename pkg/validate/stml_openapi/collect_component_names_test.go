//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-openapi
//ff:what TestCollectComponentNamesFourPaths — fetch/each/action(data-component:)/child 네 경로 수집

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestCollectComponentNamesFourPaths(t *testing.T) {
	pages := []stml.PageSpec{{
		Name: "page",
		Fetches: []stml.FetchBlock{{
			OperationID: "ListItems",
			Components:  []stml.ComponentRef{{Name: "FetchComp"}},
			Eaches: []stml.EachBlock{{
				Components: []stml.ComponentRef{{Name: "EachComp"}},
			}},
		}},
		Actions: []stml.ActionBlock{{
			OperationID: "CreateItem",
			Fields:      []stml.FieldBind{{Name: "StartAt", Tag: "data-component:ActionComp"}},
		}},
		Children: []stml.ChildNode{{
			Kind:      "component",
			Component: &stml.ComponentRef{Name: "ChildComp"},
		}},
	}}

	names := collectComponentNames(pages)
	for _, want := range []string{"FetchComp", "EachComp", "ActionComp", "ChildComp"} {
		if _, ok := names[want]; !ok {
			t.Errorf("missing component name %q", want)
		}
	}
	if len(names) != 4 {
		t.Errorf("expected 4 names, got %d: %+v", len(names), names)
	}
}

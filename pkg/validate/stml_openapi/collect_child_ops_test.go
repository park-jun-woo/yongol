//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-openapi
//ff:what TestCollectChildOps — ChildNode action/fetch/state/each/static에서 operationId 재귀 수집
package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestCollectChildOps(t *testing.T) {
	out := map[string]struct{}{}

	// All nil → no-op.
	collectChildOps(stml.ChildNode{Kind: "static"}, out)
	if len(out) != 0 {
		t.Fatalf("expected empty, got %+v", out)
	}

	// Action branch.
	collectChildOps(stml.ChildNode{Kind: "action", Action: &stml.ActionBlock{OperationID: "CreateItem"}}, out)
	// Fetch branch.
	collectChildOps(stml.ChildNode{Kind: "fetch", Fetch: &stml.FetchBlock{OperationID: "ListItems"}}, out)
	// State branch with nested action child.
	collectChildOps(stml.ChildNode{Kind: "state", State: &stml.StateBind{
		Condition: "empty",
		Children: []stml.ChildNode{
			{Kind: "action", Action: &stml.ActionBlock{OperationID: "RetryItem"}},
		},
	}}, out)
	// Each branch with a row action (page-flow Phase006).
	collectChildOps(stml.ChildNode{Kind: "each", Each: &stml.EachBlock{
		Field: "photos",
		Children: []stml.ChildNode{
			{Kind: "action", Action: &stml.ActionBlock{OperationID: "DeletePhoto"}},
		},
	}}, out)
	// Static branch with a nested action.
	collectChildOps(stml.ChildNode{Kind: "static", Static: &stml.StaticElement{
		Children: []stml.ChildNode{
			{Kind: "action", Action: &stml.ActionBlock{OperationID: "StarPhoto"}},
		},
	}}, out)

	for _, id := range []string{"CreateItem", "ListItems", "RetryItem", "DeletePhoto", "StarPhoto"} {
		if _, ok := out[id]; !ok {
			t.Errorf("missing operationId %q", id)
		}
	}
}

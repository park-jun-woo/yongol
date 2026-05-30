//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestCollectChildOps — ChildNode action/fetch/state에서 operationId 재귀 수집

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

	for _, id := range []string{"CreateItem", "ListItems", "RetryItem"} {
		if _, ok := out[id]; !ok {
			t.Errorf("missing operationId %q", id)
		}
	}
}

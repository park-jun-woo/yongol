//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestCollectFetchOps — FetchBlock + nested fetches + state children operationId 재귀 수집

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestCollectFetchOps(t *testing.T) {
	out := map[string]struct{}{}
	f := stml.FetchBlock{
		OperationID:   "ListItems",
		NestedFetches: []stml.FetchBlock{{OperationID: "ListSub"}},
		Children: []stml.ChildNode{
			{Kind: "action", Action: &stml.ActionBlock{OperationID: "CreateItem"}},
		},
	}
	collectFetchOps(f, out)
	for _, id := range []string{"ListItems", "ListSub", "CreateItem"} {
		if _, ok := out[id]; !ok {
			t.Errorf("missing operationId %q", id)
		}
	}
}

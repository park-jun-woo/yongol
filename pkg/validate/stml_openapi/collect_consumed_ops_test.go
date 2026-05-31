//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-openapi
//ff:what TestCollectConsumedOps — 전 페이지 data-fetch/data-action operationId 집합 수집
package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestCollectConsumedOps(t *testing.T) {
	// Empty input.
	if got := collectConsumedOps(nil); len(got) != 0 {
		t.Fatalf("expected empty, got %+v", got)
	}

	pages := []stml.PageSpec{{
		Name:    "page",
		Fetches: []stml.FetchBlock{{OperationID: "ListItems"}},
		Actions: []stml.ActionBlock{{OperationID: "CreateItem"}},
	}}
	out := collectConsumedOps(pages)
	for _, id := range []string{"ListItems", "CreateItem"} {
		if _, ok := out[id]; !ok {
			t.Errorf("missing operationId %q", id)
		}
	}
}

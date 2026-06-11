//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-openapi
//ff:what collectPageFetchOps — 중첩 포함 모든 fetch operationId 수집 검증

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestCollectPageFetchOps(t *testing.T) {
	page := stml.PageSpec{
		Fetches: []stml.FetchBlock{
			{OperationID: "GetRule", NestedFetches: []stml.FetchBlock{{OperationID: "GetNested"}}},
			{OperationID: "ListItems"},
		},
	}
	ops := collectPageFetchOps(page)
	for _, want := range []string{"GetRule", "GetNested", "ListItems"} {
		if !ops[want] {
			t.Errorf("missing %q in %+v", want, ops)
		}
	}
	if ops["NoSuch"] {
		t.Errorf("unexpected op present")
	}

	if got := collectPageFetchOps(stml.PageSpec{}); len(got) != 0 {
		t.Errorf("no fetches should yield empty set, got %+v", got)
	}
}

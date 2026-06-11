//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what collectEachRowActions — fetch/static/state/link 중첩 아래 each 행 액션 재귀 수집 검증

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestCollectEachRowActions(t *testing.T) {
	children := []stml.ChildNode{
		{Kind: "fetch", Fetch: &stml.FetchBlock{OperationID: "ListBuildings", Children: []stml.ChildNode{
			{Kind: "static", Static: &stml.StaticElement{Children: []stml.ChildNode{
				{Kind: "each", Each: &stml.EachBlock{
					Field:   "items",
					Actions: []stml.ActionBlock{{OperationID: "DeleteBuilding", Redirect: "building-list"}},
					Children: []stml.ChildNode{
						{Kind: "each", Each: &stml.EachBlock{
							Field:   "units",
							Actions: []stml.ActionBlock{{OperationID: "DeleteUnit"}},
						}},
					},
				}},
			}}},
		}}},
	}
	var got []stml.ActionBlock
	collectEachRowActions(children, &got)
	if len(got) != 2 {
		t.Fatalf("expected 2 row actions, got %d: %+v", len(got), got)
	}
	if got[0].OperationID != "DeleteBuilding" || got[1].OperationID != "DeleteUnit" {
		t.Errorf("row actions = %q, %q — want DeleteBuilding then DeleteUnit", got[0].OperationID, got[1].OperationID)
	}

	var rest []stml.ActionBlock
	collectEachRowActions(nil, &rest)
	if rest != nil {
		t.Errorf("expected nil for no children, got %+v", rest)
	}
}

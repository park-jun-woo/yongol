//ff:func feature=stml-gen type=test control=sequence
//ff:what populateRowActionArgs — fetch 컨텍스트 해석·static/state 중첩·중첩 fetch·비-item 액션 스킵 검증
package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestPopulateRowActionArgs(t *testing.T) {
	itemTypes := map[string]map[string]map[string]string{
		"GetUnit":    {"photos": {"id": "integer"}},
		"ListExtras": {"extras": {"key": "string"}},
	}
	ppt := map[string]map[string]string{
		"DeletePhoto": {"photoId": "integer"},
		"DropExtra":   {"extraKey": "integer"},
	}

	rowAction := stmlparser.ActionBlock{
		OperationID: "DeletePhoto",
		Params:      []stmlparser.ParamBind{{Name: "photoId", Source: "item.id"}},
	}
	plainAction := stmlparser.ActionBlock{
		OperationID: "Refresh",
		Params:      []stmlparser.ParamBind{{Name: "unitId", Source: "route.UnitID"}},
	}
	each := stmlparser.EachBlock{Field: "photos"}
	each.Children = []stmlparser.ChildNode{
		{Kind: "action", Action: &rowAction},
		{Kind: "static", Static: &stmlparser.StaticElement{Children: []stmlparser.ChildNode{
			{Kind: "action", Action: &plainAction},
		}}},
	}

	nestedAction := stmlparser.ActionBlock{
		OperationID: "DropExtra",
		Params:      []stmlparser.ParamBind{{Name: "extraKey", Source: "item.key"}},
	}
	nestedEach := stmlparser.EachBlock{Field: "extras"}
	nestedEach.Children = []stmlparser.ChildNode{
		{Kind: "state", State: &stmlparser.StateBind{Children: []stmlparser.ChildNode{
			{Kind: "action", Action: &nestedAction},
		}}},
	}
	nestedFetch := stmlparser.FetchBlock{
		OperationID: "ListExtras",
		Children:    []stmlparser.ChildNode{{Kind: "each", Each: &nestedEach}},
	}

	fetch := stmlparser.FetchBlock{
		OperationID: "GetUnit",
		Eaches:      []stmlparser.EachBlock{each},
		Children: []stmlparser.ChildNode{
			{Kind: "each", Each: &each},
			{Kind: "fetch", Fetch: &nestedFetch},
		},
	}
	page := stmlparser.PageSpec{
		Fetches:  []stmlparser.FetchBlock{fetch},
		Children: []stmlparser.ChildNode{{Kind: "fetch", Fetch: &fetch}},
	}

	populateRowActionArgs(&page, itemTypes, ppt)

	// integer item field → no Number wrapping.
	if rowAction.RowMutateArg != "{ photoId: item.id }" {
		t.Errorf("rowAction.RowMutateArg = %q", rowAction.RowMutateArg)
	}
	// string item field bound to an integer path param → Number wrapping.
	if nestedAction.RowMutateArg != "{ extraKey: Number(item.key) }" {
		t.Errorf("nestedAction.RowMutateArg = %q", nestedAction.RowMutateArg)
	}
	// actions without item params keep the hoisted-closure path.
	if plainAction.RowMutateArg != "" {
		t.Errorf("plainAction.RowMutateArg = %q, want empty", plainAction.RowMutateArg)
	}
}

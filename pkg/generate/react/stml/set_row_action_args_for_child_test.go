//ff:func feature=stml-gen type=test control=sequence
//ff:what setRowActionArgsForChild — each/fetch/static/state/기타 kind 분기 검증
package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestSetRowActionArgsForChild(t *testing.T) {
	itemTypes := map[string]map[string]map[string]string{"GetUnit": {"photos": {"id": "integer"}}}
	mkAction := func() *stmlparser.ActionBlock {
		return &stmlparser.ActionBlock{
			OperationID: "DeletePhoto",
			Params:      []stmlparser.ParamBind{{Name: "photoId", Source: "item.id"}},
		}
	}

	// each kind
	a1 := mkAction()
	each := stmlparser.EachBlock{Field: "photos", Children: []stmlparser.ChildNode{{Kind: "action", Action: a1}}}
	ch := stmlparser.ChildNode{Kind: "each", Each: &each}
	setRowActionArgsForChild(&ch, "GetUnit", itemTypes, nil)
	if a1.RowMutateArg != "{ photoId: item.id }" {
		t.Errorf("each kind: %q", a1.RowMutateArg)
	}

	// fetch kind resolves the nested operationId
	a2 := mkAction()
	each2 := stmlparser.EachBlock{Field: "photos", Children: []stmlparser.ChildNode{{Kind: "action", Action: a2}}}
	fetch := stmlparser.FetchBlock{OperationID: "GetUnit", Children: []stmlparser.ChildNode{{Kind: "each", Each: &each2}}}
	ch = stmlparser.ChildNode{Kind: "fetch", Fetch: &fetch}
	setRowActionArgsForChild(&ch, "", itemTypes, nil)
	if a2.RowMutateArg != "{ photoId: item.id }" {
		t.Errorf("fetch kind: %q", a2.RowMutateArg)
	}

	// static / state recursion (nil pointers are tolerated)
	a3 := mkAction()
	each3 := stmlparser.EachBlock{Field: "photos", Children: []stmlparser.ChildNode{{Kind: "action", Action: a3}}}
	st := stmlparser.StaticElement{Children: []stmlparser.ChildNode{{Kind: "each", Each: &each3}}}
	ch = stmlparser.ChildNode{Kind: "static", Static: &st}
	setRowActionArgsForChild(&ch, "GetUnit", itemTypes, nil)
	if a3.RowMutateArg != "{ photoId: item.id }" {
		t.Errorf("static kind: %q", a3.RowMutateArg)
	}
	ch = stmlparser.ChildNode{Kind: "static"}
	setRowActionArgsForChild(&ch, "GetUnit", itemTypes, nil) // no panic

	a4 := mkAction()
	each4 := stmlparser.EachBlock{Field: "photos", Children: []stmlparser.ChildNode{{Kind: "action", Action: a4}}}
	sb := stmlparser.StateBind{Children: []stmlparser.ChildNode{{Kind: "each", Each: &each4}}}
	ch = stmlparser.ChildNode{Kind: "state", State: &sb}
	setRowActionArgsForChild(&ch, "GetUnit", itemTypes, nil)
	if a4.RowMutateArg != "{ photoId: item.id }" {
		t.Errorf("state kind: %q", a4.RowMutateArg)
	}
	ch = stmlparser.ChildNode{Kind: "state"}
	setRowActionArgsForChild(&ch, "GetUnit", itemTypes, nil) // no panic

	// unrelated kind is ignored
	ch = stmlparser.ChildNode{Kind: "bind", Bind: &stmlparser.FieldBind{Name: "x"}}
	setRowActionArgsForChild(&ch, "GetUnit", itemTypes, nil)
}

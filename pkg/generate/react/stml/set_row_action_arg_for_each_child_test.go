//ff:func feature=stml-gen type=test control=sequence
//ff:what setRowActionArgForEachChild — action(item 유/무)·static·state·nil 분기 검증
package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestSetRowActionArgForEachChild(t *testing.T) {
	types := map[string]string{"id": "integer"}

	// action with item param → arg recorded
	a := stmlparser.ActionBlock{
		OperationID: "DeletePhoto",
		Params:      []stmlparser.ParamBind{{Name: "photoId", Source: "item.id"}},
	}
	ch := stmlparser.ChildNode{Kind: "action", Action: &a}
	setRowActionArgForEachChild(&ch, types, nil)
	if a.RowMutateArg != "{ photoId: item.id }" {
		t.Errorf("item action: %q", a.RowMutateArg)
	}

	// action without item params keeps the hoisted-closure path
	b := stmlparser.ActionBlock{
		OperationID: "Refresh",
		Params:      []stmlparser.ParamBind{{Name: "unitId", Source: "route.UnitID"}},
	}
	ch = stmlparser.ChildNode{Kind: "action", Action: &b}
	setRowActionArgForEachChild(&ch, types, nil)
	if b.RowMutateArg != "" {
		t.Errorf("route-only action must stay empty: %q", b.RowMutateArg)
	}

	// static / state recursion + nil tolerance
	c := stmlparser.ActionBlock{
		OperationID: "StarPhoto",
		Params:      []stmlparser.ParamBind{{Name: "photoId", Source: "item.id"}},
	}
	st := stmlparser.StaticElement{Children: []stmlparser.ChildNode{{Kind: "action", Action: &c}}}
	ch = stmlparser.ChildNode{Kind: "static", Static: &st}
	setRowActionArgForEachChild(&ch, types, nil)
	if c.RowMutateArg != "{ photoId: item.id }" {
		t.Errorf("static-nested action: %q", c.RowMutateArg)
	}
	ch = stmlparser.ChildNode{Kind: "static"}
	setRowActionArgForEachChild(&ch, types, nil) // no panic

	d := stmlparser.ActionBlock{
		OperationID: "HidePhoto",
		Params:      []stmlparser.ParamBind{{Name: "photoId", Source: "item.id"}},
	}
	sb := stmlparser.StateBind{Children: []stmlparser.ChildNode{{Kind: "action", Action: &d}}}
	ch = stmlparser.ChildNode{Kind: "state", State: &sb}
	setRowActionArgForEachChild(&ch, types, nil)
	if d.RowMutateArg != "{ photoId: item.id }" {
		t.Errorf("state-nested action: %q", d.RowMutateArg)
	}
	ch = stmlparser.ChildNode{Kind: "state"}
	setRowActionArgForEachChild(&ch, types, nil) // no panic

	// unrelated kind ignored
	ch = stmlparser.ChildNode{Kind: "bind", Bind: &stmlparser.FieldBind{Name: "x"}}
	setRowActionArgForEachChild(&ch, types, nil)
}

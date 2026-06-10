//ff:func feature=stml-gen type=test control=sequence
//ff:what setRowActionArgsInEachChildren — 자식 슬라이스 순회 위임 검증
package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestSetRowActionArgsInEachChildren(t *testing.T) {
	a := stmlparser.ActionBlock{
		OperationID: "DeletePhoto",
		Params:      []stmlparser.ParamBind{{Name: "photoId", Source: "item.id"}},
	}
	b := stmlparser.ActionBlock{
		OperationID: "StarPhoto",
		Params:      []stmlparser.ParamBind{{Name: "photoId", Source: "item.id"}},
	}
	children := []stmlparser.ChildNode{
		{Kind: "action", Action: &a},
		{Kind: "action", Action: &b},
	}
	setRowActionArgsInEachChildren(children, map[string]string{"id": "integer"}, nil)
	if a.RowMutateArg != "{ photoId: item.id }" || b.RowMutateArg != "{ photoId: item.id }" {
		t.Errorf("RowMutateArgs = %q, %q", a.RowMutateArg, b.RowMutateArg)
	}
}

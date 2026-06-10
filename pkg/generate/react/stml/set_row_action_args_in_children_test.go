//ff:func feature=stml-gen type=test control=sequence
//ff:what setRowActionArgsInChildren — 슬라이스 순회 위임 검증
package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestSetRowActionArgsInChildren(t *testing.T) {
	a := stmlparser.ActionBlock{
		OperationID: "DeletePhoto",
		Params:      []stmlparser.ParamBind{{Name: "photoId", Source: "item.id"}},
	}
	each := stmlparser.EachBlock{Field: "photos", Children: []stmlparser.ChildNode{
		{Kind: "action", Action: &a},
	}}
	children := []stmlparser.ChildNode{{Kind: "each", Each: &each}}
	itemTypes := map[string]map[string]map[string]string{"GetUnit": {"photos": {"id": "integer"}}}

	setRowActionArgsInChildren(children, "GetUnit", itemTypes, nil)
	if a.RowMutateArg != "{ photoId: item.id }" {
		t.Errorf("RowMutateArg = %q", a.RowMutateArg)
	}
}

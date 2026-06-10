//ff:func feature=stml-gen type=test control=sequence
//ff:what setRowActionArgsInEach — item 스키마 해석 성공/실패(미상 op) 분기 검증
package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestSetRowActionArgsInEach(t *testing.T) {
	itemTypes := map[string]map[string]map[string]string{"GetUnit": {"photos": {"id": "integer"}}}
	ppt := map[string]map[string]string{"DeletePhoto": {"photoId": "integer"}}

	// known op + field → integer item type, no Number wrapping
	a := stmlparser.ActionBlock{
		OperationID: "DeletePhoto",
		Params:      []stmlparser.ParamBind{{Name: "photoId", Source: "item.id"}},
	}
	e := stmlparser.EachBlock{Field: "photos", Children: []stmlparser.ChildNode{{Kind: "action", Action: &a}}}
	setRowActionArgsInEach(&e, "GetUnit", itemTypes, ppt)
	if a.RowMutateArg != "{ photoId: item.id }" {
		t.Errorf("known schema: %q", a.RowMutateArg)
	}

	// unknown op → nil item schema → integer path param gets Number wrapping
	b := stmlparser.ActionBlock{
		OperationID: "DeletePhoto",
		Params:      []stmlparser.ParamBind{{Name: "photoId", Source: "item.id"}},
	}
	e2 := stmlparser.EachBlock{Field: "photos", Children: []stmlparser.ChildNode{{Kind: "action", Action: &b}}}
	setRowActionArgsInEach(&e2, "UnknownOp", itemTypes, ppt)
	if b.RowMutateArg != "{ photoId: Number(item.id) }" {
		t.Errorf("unknown schema: %q", b.RowMutateArg)
	}
}

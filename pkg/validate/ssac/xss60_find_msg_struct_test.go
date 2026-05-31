//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what xss60BuildTableMap/xss60ResolveVarModel/xss60FindMsgStruct 단위 검증
package ssac

import (
	"testing"

	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestXss60FindMsgStruct(t *testing.T) {
	fn := parsessac.ServiceFunc{
		Param: &parsessac.ParamInfo{TypeName: "OnOrderCompletedMessage"},
		Structs: []parsessac.StructInfo{
			{Name: "Other"},
			{Name: "OnOrderCompletedMessage"},
		},
	}
	got := xss60FindMsgStruct(fn)
	if got == nil || got.Name != "OnOrderCompletedMessage" {
		t.Errorf("findMsgStruct = %v", got)
	}

	// no matching struct -> nil
	fn2 := parsessac.ServiceFunc{
		Param:   &parsessac.ParamInfo{TypeName: "Missing"},
		Structs: []parsessac.StructInfo{{Name: "Other"}},
	}
	if got := xss60FindMsgStruct(fn2); got != nil {
		t.Errorf("expected nil, got %v", got)
	}
}

//ff:func feature=stml-gen type=test control=sequence
//ff:what required integer route 파라미터 useQuery가 enabled 가드 없이 Number(id)를 유지하는지 검증 (BUG-136)
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestRenderUseQuery_RequiredParamNoEnabledGuard(t *testing.T) {
	// A required param (fetch-consumed ":id") is always present — no enabled
	// guard, byte-identical to the pre-BUG-136 output.
	f := stmlparser.FetchBlock{
		OperationID: "GetBuilding",
		Params:      []stmlparser.ParamBind{{Name: "id", Source: "route.id"}},
	}
	ppt := map[string]map[string]string{"GetBuilding": {"id": "integer"}}

	code := renderUseQuery(f, ppt)
	if !strings.Contains(code, "id: Number(id)") {
		t.Errorf("required param should stay Number(id):\n%s", code)
	}
	if strings.Contains(code, "enabled:") {
		t.Errorf("required param query should not emit enabled:\n%s", code)
	}
}

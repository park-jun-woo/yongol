//ff:func feature=stml-gen type=test control=sequence
//ff:what optional integer route 파라미터 useQuery가 enabled 가드와 null 가드된 Number()를 방출하는지 검증 (BUG-136)
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestRenderUseQuery_OptionalParamEnabledGuard(t *testing.T) {
	// A fetch whose integer param is optional (the rare explicit data-route
	// ":Id?" case): the query must gate on the param and guard the Number().
	f := stmlparser.FetchBlock{
		OperationID: "GetBuilding",
		Params:      []stmlparser.ParamBind{{Name: "id", Source: "route.id", Optional: true}},
	}
	ppt := map[string]map[string]string{"GetBuilding": {"id": "integer"}}

	code := renderUseQuery(f, ppt)
	// BUG-137: the arg stays plain Number(id) (type number); the empty-value
	// call is blocked by the enabled guard below, not by mutilating the arg.
	if !strings.Contains(code, "id: Number(id) }") {
		t.Errorf("optional param arg should be plain Number(id):\n%s", code)
	}
	if strings.Contains(code, "id != null ? Number(id) : undefined") {
		t.Errorf("optional param arg should no longer be null-guarded:\n%s", code)
	}
	if !strings.Contains(code, "enabled: Number.isFinite(Number(id)),") {
		t.Errorf("optional integer query missing enabled guard:\n%s", code)
	}
}

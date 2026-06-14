//ff:func feature=stml-gen type=test control=sequence
//ff:what optional string route 파라미터 useQuery가 truthiness enabled 가드를 쓰고 Number() 래핑하지 않는지 검증 (BUG-136)
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestRenderUseQuery_OptionalStringParamEnabledGuard(t *testing.T) {
	// A non-integer optional param gates on truthiness and is not Number-wrapped.
	f := stmlparser.FetchBlock{
		OperationID: "GetUser",
		Params:      []stmlparser.ParamBind{{Name: "slug", Source: "route.slug", Optional: true}},
	}
	ppt := map[string]map[string]string{"GetUser": {"slug": "string"}}

	code := renderUseQuery(f, ppt)
	if !strings.Contains(code, "enabled: !!slug,") {
		t.Errorf("optional string query missing truthiness guard:\n%s", code)
	}
	if strings.Contains(code, "Number(slug)") {
		t.Errorf("string param should not be Number-wrapped:\n%s", code)
	}
}

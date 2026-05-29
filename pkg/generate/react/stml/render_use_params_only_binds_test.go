//ff:func feature=stml-gen type=test control=sequence
//ff:what renderUseParamsWithRoute ParamBind만 있을 때 검증

package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestRenderUseParamsWithRoute_OnlyParamBinds(t *testing.T) {
	params := []stmlparser.ParamBind{{Name: "id", Source: "route.id"}}
	got := renderUseParamsWithRoute(params, "")
	want := "const { id } = useParams()"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

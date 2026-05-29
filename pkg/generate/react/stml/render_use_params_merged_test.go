//ff:func feature=stml-gen type=test control=sequence
//ff:what renderUseParamsWithRoute ParamBind + 라우트 패턴 병합 검증

package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestRenderUseParamsWithRoute_MergedParams(t *testing.T) {
	params := []stmlparser.ParamBind{
		{Name: "buildingId", Source: "route.buildingId"},
		{Name: "id", Source: "route.id"},
	}
	got := renderUseParamsWithRoute(params, "/buildings/:buildingId/units/:id")
	want := "const { buildingId, id } = useParams()"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

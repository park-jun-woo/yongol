//ff:func feature=stml-gen type=test control=sequence
//ff:what renderUseParamsWithRoute 라우트 패턴이 추가 파라미터를 보충하는 경우 검증

package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestRenderUseParamsWithRoute_RouteAddsExtra(t *testing.T) {
	params := []stmlparser.ParamBind{{Name: "id", Source: "route.id"}}
	got := renderUseParamsWithRoute(params, "/buildings/:buildingId/units/:id")
	want := "const { id, buildingId } = useParams()"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

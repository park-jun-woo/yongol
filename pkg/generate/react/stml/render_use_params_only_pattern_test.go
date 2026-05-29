//ff:func feature=stml-gen type=test control=sequence
//ff:what renderUseParamsWithRoute 라우트 패턴만 있을 때 검증

package stml

import "testing"

func TestRenderUseParamsWithRoute_OnlyRoutePattern(t *testing.T) {
	got := renderUseParamsWithRoute(nil, "/buildings/:buildingId/units/:id")
	want := "const { buildingId, id } = useParams()"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

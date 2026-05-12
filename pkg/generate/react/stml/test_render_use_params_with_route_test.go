//ff:func feature=stml-gen type=test control=sequence
//ff:what data-route 기반 다중 파라미터 useParams 생성 테스트

package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestRenderUseParamsWithRoute_NoRouteNoParams(t *testing.T) {
	got := renderUseParamsWithRoute(nil, "")
	if got != "" {
		t.Errorf("got %q, want empty", got)
	}
}

func TestRenderUseParamsWithRoute_OnlyParamBinds(t *testing.T) {
	params := []stmlparser.ParamBind{
		{Name: "id", Source: "route.id"},
	}
	got := renderUseParamsWithRoute(params, "")
	want := "const { id } = useParams()"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestRenderUseParamsWithRoute_OnlyRoutePattern(t *testing.T) {
	got := renderUseParamsWithRoute(nil, "/buildings/:buildingId/units/:id")
	want := "const { buildingId, id } = useParams()"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

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

func TestRenderUseParamsWithRoute_RouteAddsExtra(t *testing.T) {
	// ParamBind only has "id", but route pattern also has "buildingId"
	params := []stmlparser.ParamBind{
		{Name: "id", Source: "route.id"},
	}
	got := renderUseParamsWithRoute(params, "/buildings/:buildingId/units/:id")
	want := "const { id, buildingId } = useParams()"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

//ff:func feature=gen-react type=test control=iteration dimension=1
//ff:what navRoutePatterns — 페이지명→해석 라우트 패턴 맵 구성 검증 (data-route 우선·유도 라우트)

package react

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestNavRoutePatterns(t *testing.T) {
	pages := []stml.PageSpec{
		{Name: "dashboard", FileName: "dashboard.html"},
		{Name: "building-detail", FileName: "building-detail.html",
			Fetches: []stml.FetchBlock{{
				OperationID: "GetBuilding",
				Params:      []stml.ParamBind{{Name: "building_id", Source: "route.BuildingID"}},
			}}},
		{Name: "custom", FileName: "custom.html", Route: "/my/route"},
	}

	got := navRoutePatterns(pages)
	want := map[string]string{
		"dashboard":       "/dashboard",
		"building-detail": "/buildings/:BuildingID",
		"custom":          "/my/route",
	}
	if len(got) != len(want) {
		t.Fatalf("len = %d, want %d: %+v", len(got), len(want), got)
	}
	for k, v := range want {
		if got[k] != v {
			t.Errorf("got[%q] = %q, want %q", k, got[k], v)
		}
	}
}

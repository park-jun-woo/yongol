//ff:func feature=stml-parse type=test control=sequence
//ff:what appendRouteParams — route. 소스 필터링·빈 이름 스킵·첫 등장 우선 중복 제거 검증

package stml

import (
	"reflect"
	"testing"
)

func TestAppendRouteParams(t *testing.T) {
	t.Run("filters route. sources and keeps order", func(t *testing.T) {
		seen := map[string]bool{}
		got := appendRouteParams(nil, seen, []ParamBind{
			{Name: "q", Source: "form.q"},
			{Name: "buildingId", Source: "route.BuildingID"},
			{Name: "unitId", Source: "route.UnitID"},
		}, true)
		want := []routeParam{
			{Name: "BuildingID", Required: true},
			{Name: "UnitID", Required: true},
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("first appearance wins over later duplicates", func(t *testing.T) {
		seen := map[string]bool{}
		got := appendRouteParams(nil, seen, []ParamBind{{Name: "id", Source: "route.id"}}, true)
		got = appendRouteParams(got, seen, []ParamBind{{Name: "id", Source: "route.id"}}, false)
		want := []routeParam{{Name: "id", Required: true}}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("empty name after route. prefix is skipped", func(t *testing.T) {
		seen := map[string]bool{}
		if got := appendRouteParams(nil, seen, []ParamBind{{Name: "x", Source: "route."}}, false); got != nil {
			t.Errorf("got %v, want nil", got)
		}
	})
}

//ff:func feature=stml-parse type=test control=sequence
//ff:what appendFetchRouteParams — fetch 본체·중첩 fetch 파라미터의 필수 누적 검증

package stml

import (
	"reflect"
	"testing"
)

func TestAppendFetchRouteParams(t *testing.T) {
	f := FetchBlock{
		Params: []ParamBind{{Name: "buildingId", Source: "route.BuildingID"}},
		NestedFetches: []FetchBlock{{
			Params: []ParamBind{
				{Name: "unitId", Source: "route.UnitID"},
				{Name: "buildingId", Source: "route.BuildingID"}, // dedup
			},
		}},
	}
	got := appendFetchRouteParams(nil, map[string]bool{}, f)
	want := []routeParam{
		{Name: "BuildingID", Required: true},
		{Name: "UnitID", Required: true},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

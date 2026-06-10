//ff:func feature=stml-parse type=test control=sequence
//ff:what ConsumedRouteParams — 첫 등장 순서 이름 반환·무소비 nil 검증

package stml

import (
	"reflect"
	"testing"
)

func TestConsumedRouteParams(t *testing.T) {
	t.Run("names in first-appearance order, fetch before action", func(t *testing.T) {
		p := PageSpec{
			Fetches: []FetchBlock{{
				Params: []ParamBind{
					{Name: "buildingId", Source: "route.BuildingID"},
					{Name: "unitId", Source: "route.UnitID"},
				},
			}},
			Actions: []ActionBlock{{
				Params: []ParamBind{{Name: "photoId", Source: "route.PhotoID"}},
			}},
		}
		want := []string{"BuildingID", "UnitID", "PhotoID"}
		if got := ConsumedRouteParams(p); !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("page without route.* consumption returns nil", func(t *testing.T) {
		p := PageSpec{
			Fetches: []FetchBlock{{
				Params: []ParamBind{{Name: "q", Source: "form.q"}},
			}},
		}
		if got := ConsumedRouteParams(p); got != nil {
			t.Errorf("got %v, want nil", got)
		}
	})
}

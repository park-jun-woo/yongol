//ff:func feature=stml-parse type=test control=sequence
//ff:what collectRouteParams — 등장 순서·필수/선택 구분·중복 제거·중첩 fetch/자식 action 수집 검증

package stml

import (
	"reflect"
	"testing"
)

func TestCollectRouteParams(t *testing.T) {
	t.Run("fetch params are required, action-only params optional", func(t *testing.T) {
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
		want := []routeParam{
			{Name: "BuildingID", Required: true},
			{Name: "UnitID", Required: true},
			{Name: "PhotoID", Required: false},
		}
		if got := collectRouteParams(p); !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("param consumed by fetch and action stays required once", func(t *testing.T) {
		p := PageSpec{
			Fetches: []FetchBlock{{
				Params: []ParamBind{{Name: "id", Source: "route.id"}},
			}},
			Actions: []ActionBlock{{
				Params: []ParamBind{{Name: "id", Source: "route.id"}},
			}},
		}
		want := []routeParam{{Name: "id", Required: true}}
		if got := collectRouteParams(p); !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("nested fetch params are required", func(t *testing.T) {
		p := PageSpec{
			Fetches: []FetchBlock{{
				Params: []ParamBind{{Name: "buildingId", Source: "route.BuildingID"}},
				NestedFetches: []FetchBlock{{
					Params: []ParamBind{{Name: "unitId", Source: "route.UnitID"}},
				}},
			}},
		}
		want := []routeParam{
			{Name: "BuildingID", Required: true},
			{Name: "UnitID", Required: true},
		}
		if got := collectRouteParams(p); !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("child actions inside fetch/state/each are collected as optional", func(t *testing.T) {
		p := PageSpec{
			Children: []ChildNode{{
				Kind: "fetch",
				Fetch: &FetchBlock{
					Children: []ChildNode{{
						Kind: "state",
						State: &StateBind{
							Children: []ChildNode{{
								Kind: "action",
								Action: &ActionBlock{
									Params: []ParamBind{{Name: "certId", Source: "route.CertID"}},
								},
							}},
						},
					}},
				},
			}},
		}
		want := []routeParam{{Name: "CertID", Required: false}}
		if got := collectRouteParams(p); !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("non-route sources and empty names are skipped", func(t *testing.T) {
		p := PageSpec{
			Fetches: []FetchBlock{{
				Params: []ParamBind{
					{Name: "q", Source: "form.q"},
					{Name: "x", Source: "route."},
				},
			}},
		}
		if got := collectRouteParams(p); got != nil {
			t.Errorf("got %v, want nil", got)
		}
	})
}

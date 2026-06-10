//ff:func feature=stml-parse type=test control=sequence
//ff:what pageHasRouteParam — fetch/action route. 소스 존재/부재 검증

package stml

import "testing"

func TestPageHasRouteParam(t *testing.T) {
	t.Run("route source in fetch", func(t *testing.T) {
		p := PageSpec{
			Fetches: []FetchBlock{
				{Params: []ParamBind{{Source: "route.ID"}}},
			},
		}
		if !pageHasRouteParam(p) {
			t.Errorf("expected true for fetch route source")
		}
	})

	t.Run("route source in action only", func(t *testing.T) {
		p := PageSpec{
			Fetches: []FetchBlock{
				{Params: []ParamBind{{Source: "query.q"}}},
			},
			Actions: []ActionBlock{
				{Params: []ParamBind{{Source: "route.ReservationID"}}},
			},
		}
		if !pageHasRouteParam(p) {
			t.Errorf("expected true for action route source")
		}
	})

	t.Run("no route source", func(t *testing.T) {
		p := PageSpec{
			Fetches: []FetchBlock{
				{Params: []ParamBind{{Source: "query.q"}}},
			},
			Actions: []ActionBlock{
				{Params: []ParamBind{{Source: "form.name"}}},
			},
		}
		if pageHasRouteParam(p) {
			t.Errorf("expected false when no route source")
		}
	})
}

//ff:func feature=stml-parse type=test control=sequence
//ff:what paramsHaveRouteSource — route. 접두사 소스 존재/부재 검증

package stml

import "testing"

func TestParamsHaveRouteSource(t *testing.T) {
	t.Run("has route source", func(t *testing.T) {
		params := []ParamBind{
			{Source: "query.q"},
			{Source: "route.ID"},
		}
		if !paramsHaveRouteSource(params) {
			t.Errorf("expected true for route source")
		}
	})

	t.Run("no route source", func(t *testing.T) {
		params := []ParamBind{
			{Source: "query.q"},
			{Source: "form.name"},
		}
		if paramsHaveRouteSource(params) {
			t.Errorf("expected false without route source")
		}
	})

	t.Run("empty params", func(t *testing.T) {
		if paramsHaveRouteSource(nil) {
			t.Errorf("expected false for empty params")
		}
	})
}

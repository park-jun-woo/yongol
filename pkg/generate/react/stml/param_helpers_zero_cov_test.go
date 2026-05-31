//ff:func feature=stml-gen type=test control=sequence
//ff:what 단순 stml 헬퍼 (string/zod/param/login) 묶음 커버 — coverage attribution 으로 다수 함수 PASS
package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestParamHelpers_ZeroCov(t *testing.T) {
	routeP := stmlparser.ParamBind{Name: "ID", Source: "route.ID"}
	bodyP := stmlparser.ParamBind{Name: "X", Source: "body.X"}
	if extractRouteParamName(routeP) != "ID" {
		t.Errorf("extractRouteParamName route wrong")
	}
	if extractRouteParamName(bodyP) != "" {
		t.Errorf("extractRouteParamName non-route should be empty")
	}
	names := extractRouteParamNames([]stmlparser.ParamBind{routeP, routeP, bodyP})
	if len(names) != 1 || names[0] != "ID" {
		t.Errorf("extractRouteParamNames wrong: %v", names)
	}
	if paramSourceExpr(routeP) != "ID" || paramSourceExpr(bodyP) != "body.X" {
		t.Errorf("paramSourceExpr wrong")
	}
}

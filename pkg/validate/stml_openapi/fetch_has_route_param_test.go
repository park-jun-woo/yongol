//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what fetchHasRouteParam — route.* 소스 path 파라미터 검출, item.*·무파라미터 침묵 검증

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestFetchHasRouteParam(t *testing.T) {
	withRoute := stml.FetchBlock{Params: []stml.ParamBind{{Name: "RuleID", Source: "route.RuleID"}}}
	if !fetchHasRouteParam(withRoute) {
		t.Errorf("route.* source should be detected")
	}

	itemSource := stml.FetchBlock{Params: []stml.ParamBind{{Name: "id", Source: "item.id"}}}
	if fetchHasRouteParam(itemSource) {
		t.Errorf("item.* source is not a route path param")
	}

	if fetchHasRouteParam(stml.FetchBlock{}) {
		t.Errorf("no params should be false")
	}
}

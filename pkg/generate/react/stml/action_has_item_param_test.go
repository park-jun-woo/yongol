//ff:func feature=stml-gen type=test control=sequence
//ff:what actionHasItemParam — item.* 소스 유무 판별 검증
package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestActionHasItemParam(t *testing.T) {
	withItem := stmlparser.ActionBlock{Params: []stmlparser.ParamBind{
		{Name: "buildingId", Source: "route.BuildingID"},
		{Name: "photoId", Source: "item.id"},
	}}
	if !actionHasItemParam(withItem) {
		t.Error("expected true for item.* source")
	}
	routeOnly := stmlparser.ActionBlock{Params: []stmlparser.ParamBind{
		{Name: "buildingId", Source: "route.BuildingID"},
	}}
	if actionHasItemParam(routeOnly) {
		t.Error("expected false for route-only params")
	}
	if actionHasItemParam(stmlparser.ActionBlock{}) {
		t.Error("expected false for no params")
	}
}

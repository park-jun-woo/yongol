//ff:func feature=stml-gen type=test control=sequence
//ff:what TestRedirectUsesRouteParams — route.* 소스 존재 판별 검증
package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestRedirectUsesRouteParams(t *testing.T) {
	with := stmlparser.ActionBlock{RedirectParams: []stmlparser.LinkParamBind{
		{Source: "id", Segment: "ContractID"},
		{Source: "route.BuildingID", Segment: "BuildingID"},
	}}
	if !redirectUsesRouteParams(with) {
		t.Errorf("route source: expected true")
	}

	without := stmlparser.ActionBlock{RedirectParams: []stmlparser.LinkParamBind{
		{Source: "id", Segment: "ContractID"},
	}}
	if redirectUsesRouteParams(without) {
		t.Errorf("respField only: expected false")
	}

	if redirectUsesRouteParams(stmlparser.ActionBlock{}) {
		t.Errorf("no params: expected false")
	}
}

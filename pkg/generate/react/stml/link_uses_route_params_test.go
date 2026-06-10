//ff:func feature=stml-gen type=test control=sequence
//ff:what TestLinkUsesRouteParams — route.* 소스 유무 판별 검증
package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestLinkUsesRouteParams(t *testing.T) {
	withRoute := stmlparser.LinkRef{Params: []stmlparser.LinkParamBind{
		{Source: "item.id", Segment: "A"},
		{Source: "route.BuildingID", Segment: "B"},
	}}
	if !linkUsesRouteParams(withRoute) {
		t.Error("expected true with a route.* source")
	}
	itemOnly := stmlparser.LinkRef{Params: []stmlparser.LinkParamBind{{Source: "item.id"}}}
	if linkUsesRouteParams(itemOnly) {
		t.Error("expected false with item.* sources only")
	}
	if linkUsesRouteParams(stmlparser.LinkRef{}) {
		t.Error("expected false without params")
	}
}

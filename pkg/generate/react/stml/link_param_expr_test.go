//ff:func feature=stml-gen type=test control=sequence
//ff:what TestLinkParamExpr — route.* 변수 변환·item.* 패스스루 검증
package stml

import "testing"

func TestLinkParamExpr(t *testing.T) {
	if got := linkParamExpr("route.BuildingID"); got != "BuildingID" {
		t.Errorf("route: got %q", got)
	}
	if got := linkParamExpr("item.id"); got != "item.id" {
		t.Errorf("item: got %q", got)
	}
	if got := linkParamExpr("item.photo.id"); got != "item.photo.id" {
		t.Errorf("dotted item: got %q", got)
	}
}

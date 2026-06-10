//ff:func feature=stml-gen type=test control=sequence
//ff:what TestRedirectParamExpr — route 소스는 useParams 변수, respField는 data.<field> 변환 검증
package stml

import "testing"

func TestRedirectParamExpr(t *testing.T) {
	if got := redirectParamExpr("route.BuildingID"); got != "BuildingID" {
		t.Errorf("route source = %q, want %q", got, "BuildingID")
	}
	if got := redirectParamExpr("id"); got != "data.id" {
		t.Errorf("respField = %q, want %q", got, "data.id")
	}
}

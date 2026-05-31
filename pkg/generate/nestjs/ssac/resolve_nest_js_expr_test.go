//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestZeroCov — 0% render/util 함수 (controllerRoutePrefix / formatCallTarget / render*Op / resolveDataKey 등) 회귀
package ssac

import (
	"testing"
)

func TestResolveNestJSExpr(t *testing.T) {
	if got := resolveNestJSExpr("request.email"); got != "body.email" {
		t.Errorf("request rewrite = %q", got)
	}
	if got := resolveNestJSExpr("user.id"); got != "user.id" {
		t.Errorf("passthrough = %q", got)
	}
}

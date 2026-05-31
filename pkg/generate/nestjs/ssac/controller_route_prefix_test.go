//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestZeroCov — 0% render/util 함수 (controllerRoutePrefix / formatCallTarget / render*Op / resolveDataKey 등) 회귀
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestControllerRoutePrefix(t *testing.T) {
	if got := controllerRoutePrefix(&ir.ServicePlan{URLPath: "/courses/:id"}); got != "courses" {
		t.Errorf("URLPath = %q, want courses", got)
	}
	if got := controllerRoutePrefix(&ir.ServicePlan{URLPath: "/courses"}); got != "courses" {
		t.Errorf("single-seg = %q, want courses", got)
	}
	got := controllerRoutePrefix(&ir.ServicePlan{URLPath: "", Feature: "Auth"})
	if got != "auth" {
		t.Errorf("empty path fallback = %q, want auth", got)
	}
}

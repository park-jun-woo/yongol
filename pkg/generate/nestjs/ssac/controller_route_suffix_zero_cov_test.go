//ff:func feature=gen-nestjs type=test control=sequence
//ff:what nestjs/ssac 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestControllerRouteSuffix_ZeroCov(t *testing.T) {
	plan := &ir.ServicePlan{URLPath: "/courses/{id}/enroll"}
	if got := controllerRouteSuffix(plan); got != ":id/enroll" {
		t.Errorf("suffix=%q", got)
	}
	if got := controllerRouteSuffix(&ir.ServicePlan{URLPath: "/courses"}); got != "" {
		t.Errorf("single segment should be empty, got %q", got)
	}
	if got := controllerRouteSuffix(&ir.ServicePlan{}); got != "" {
		t.Errorf("empty path should be empty")
	}
}

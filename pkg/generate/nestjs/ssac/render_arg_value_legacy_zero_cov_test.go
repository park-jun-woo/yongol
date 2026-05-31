//ff:func feature=gen-nestjs type=test control=sequence
//ff:what nestjs/ssac 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderArgValueLegacy_ZeroCov(t *testing.T) {
	if got := renderArgValueLegacy(ir.FieldArg{Source: "request"}, ""); got != "params" {
		t.Errorf("request empty col: %q", got)
	}
	if got := renderArgValueLegacy(ir.FieldArg{Source: "currentUser"}, ""); got != "user" {
		t.Errorf("currentUser empty col: %q", got)
	}
	if got := renderArgValueLegacy(ir.FieldArg{Source: "request"}, "name"); got != "body.name" {
		t.Errorf("request col: %q", got)
	}
	if got := renderArgValueLegacy(ir.FieldArg{Source: "var"}, "x"); got != "var.x" {
		t.Errorf("var col: %q", got)
	}
}

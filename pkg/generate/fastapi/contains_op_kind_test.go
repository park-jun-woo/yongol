//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestFastapiHelpers — fastapi plan/package/route 헬퍼 검증 (Op 종류·외부 패키지 수집·라우트 해석)
package fastapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestContainsOpKind(t *testing.T) {
	ops := []ir.Op{{Kind: ir.OpCall}, {Kind: ir.OpAuth}}
	if !containsOpKind(ops, ir.OpAuth) {
		t.Error("expected auth op present")
	}
	if containsOpKind(ops, ir.OpPublish) {
		t.Error("publish op should be absent")
	}
	if containsOpKind(nil, ir.OpAuth) {
		t.Error("nil ops should be false")
	}
}

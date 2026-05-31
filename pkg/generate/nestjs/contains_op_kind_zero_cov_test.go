//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestNestifyPath_ZeroCov — {param} → :param 변환
package nestjs

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestContainsOpKind_ZeroCov(t *testing.T) {
	ops := []ir.Op{{Kind: ir.OpAuth}, {Kind: ir.OpGet}}
	if !containsOpKind(ops, ir.OpAuth) {
		t.Error("expected OpAuth present")
	}
	if containsOpKind(ops, ir.OpPublish) {
		t.Error("OpPublish should be absent")
	}
}

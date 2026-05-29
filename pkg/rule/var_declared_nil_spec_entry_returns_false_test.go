//ff:func feature=rule type=test control=sequence
//ff:what TestVarDeclared_NilSpecEntry_ReturnsFalse — nil *VarDeclaredSpec 엔트리는 검사 생략

package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/pkg/toulmin"
)

func TestVarDeclared_NilSpecEntry_ReturnsFalse(t *testing.T) {
	ctx := toulmin.NewContext()
	var nilSpec *VarDeclaredSpec
	ok, ev := VarDeclared(ctx, toulmin.Specs{nilSpec})
	if ok || ev != nil {
		t.Fatalf("VarDeclared(nil spec) = (%v, %v); want (false, nil)", ok, ev)
	}
}

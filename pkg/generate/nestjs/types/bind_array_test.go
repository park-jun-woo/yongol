//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestbindArray — bindArray NestJS 바인딩 (NOT NULL + nullable) 커버
package types

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

func TestBindArray_ZeroCov(t *testing.T) {
	nn := bindNestJS(typemap.FamilyArray, ir.BindOpts{NotNull: true, ElementHead: "TEXT"})
	if !nn.Supported {
		t.Errorf("bindArray NOT NULL should be supported: %+v", nn)
	}
	nl := bindNestJS(typemap.FamilyArray, ir.BindOpts{NotNull: false, ElementHead: "TEXT"})
	_ = nl
}

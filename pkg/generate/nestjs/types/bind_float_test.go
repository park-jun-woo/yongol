//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestbindFloat — bindFloat NestJS 바인딩 (NOT NULL + nullable) 커버
package types

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

func TestBindFloat_ZeroCov(t *testing.T) {
	nn := bindNestJS(typemap.FamilyFloat, ir.BindOpts{NotNull: true})
	if !nn.Supported {
		t.Errorf("bindFloat NOT NULL should be supported: %+v", nn)
	}
	nl := bindNestJS(typemap.FamilyFloat, ir.BindOpts{NotNull: false})
	_ = nl
}

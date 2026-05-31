//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestbindEnum — bindEnum NestJS 바인딩 (NOT NULL + nullable) 커버
package types

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

func TestBindEnum_ZeroCov(t *testing.T) {
	nn := bindNestJS(typemap.FamilyEnum, ir.BindOpts{NotNull: true})
	if !nn.Supported {
		t.Errorf("bindEnum NOT NULL should be supported: %+v", nn)
	}
	nl := bindNestJS(typemap.FamilyEnum, ir.BindOpts{NotNull: false})
	_ = nl
}

//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestbindJSONB — bindJSONB NestJS 바인딩 (NOT NULL + nullable) 커버
package types

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

func TestBindJSONB_ZeroCov(t *testing.T) {
	nn := bindNestJS(typemap.FamilyJSONB, ir.BindOpts{NotNull: true})
	if !nn.Supported {
		t.Errorf("bindJSONB NOT NULL should be supported: %+v", nn)
	}
	nl := bindNestJS(typemap.FamilyJSONB, ir.BindOpts{NotNull: false})
	_ = nl
}

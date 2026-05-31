//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestNewRegistry — NewRegistry 팩토리 + Bind 메서드 커버
package types

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

func TestNewRegistry_ZeroCov(t *testing.T) {
	r := NewRegistry()
	if r == nil {
		t.Fatal("NewRegistry returned nil")
	}
	b := r.Bind(typemap.FamilyInteger, ir.BindOpts{NotNull: true})
	if !b.Supported {
		t.Errorf("integer bind should be supported")
	}
}

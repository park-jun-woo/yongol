//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestRegistry — NewRegistry 팩토리 + Bind 디스패치 검증

package types

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

func TestNewRegistry(t *testing.T) {
	r := NewRegistry()
	if r == nil {
		t.Fatal("NewRegistry() returned nil")
	}
}

func TestRegistryBind(t *testing.T) {
	r := NewRegistry()
	b := r.Bind(typemap.FamilyInteger, ir.BindOpts{NotNull: true})
	if b.Family != typemap.FamilyInteger || !b.Supported {
		t.Errorf("Bind(Integer) = %+v", b)
	}
	if b.APIType != "int" {
		t.Errorf("APIType = %q, want int", b.APIType)
	}
}

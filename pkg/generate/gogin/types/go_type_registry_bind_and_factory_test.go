//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestZeroCov — 0% util/adapter 함수 (bindGoType / GoTypeOf / registry / family dispatch / pgtype) 회귀
package types

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

func TestGoTypeRegistry_BindAndFactory(t *testing.T) {
	r := NewGoTypeRegistry()
	if r == nil {
		t.Fatal("NewGoTypeRegistry returned nil")
	}
	tb := r.Bind(typemap.FamilyInteger, ir.BindOpts{NotNull: true})
	if !tb.Supported {
		t.Errorf("integer NOT NULL should be supported")
	}
	if tb.Family != typemap.FamilyInteger {
		t.Errorf("Family = %v, want FamilyInteger", tb.Family)
	}
	if tb.DBType != "int64" {
		t.Errorf("DBType = %q, want int64", tb.DBType)
	}
}

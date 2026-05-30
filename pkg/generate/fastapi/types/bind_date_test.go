//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestBindDate — FamilyDate → SQLAlchemy Date / Python date 바인딩

package types

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

func TestBindDate(t *testing.T) {
	t.Run("NotNull", func(t *testing.T) {
		b := bindDate(ir.BindOpts{NotNull: true, DefaultLiteral: "CURRENT_DATE"})
		if b.Family != typemap.FamilyDate || !b.Supported {
			t.Errorf("unexpected binding: %+v", b)
		}
		if b.DBType != "Date" {
			t.Errorf("DBType = %q", b.DBType)
		}
		if b.APIType != "date" {
			t.Errorf("APIType = %q, want date", b.APIType)
		}
		if b.NilCheckExpr != "" {
			t.Errorf("NilCheckExpr = %q, want empty", b.NilCheckExpr)
		}
		if b.DefaultLiteral != "CURRENT_DATE" {
			t.Errorf("DefaultLiteral = %q", b.DefaultLiteral)
		}
	})
	t.Run("Nullable", func(t *testing.T) {
		b := bindDate(ir.BindOpts{NotNull: false})
		if b.APIType != "date | None" {
			t.Errorf("APIType = %q, want date | None", b.APIType)
		}
		if b.NilCheckExpr != "{var} is None" {
			t.Errorf("NilCheckExpr = %q", b.NilCheckExpr)
		}
	})
}

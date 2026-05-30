//ff:func feature=gen-fastapi type=test control=selection
//ff:what TestBindArray — PG 배열 → SQLAlchemy ARRAY/Python list 바인딩 (지원/미지원 element)

package types

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

func TestBindArray(t *testing.T) {
	t.Run("SupportedNotNull", func(t *testing.T) {
		b := bindArray(ir.BindOpts{ElementHead: "BIGINT", NotNull: true})
		if !b.Supported {
			t.Fatal("expected Supported true")
		}
		if b.Family != typemap.FamilyArray {
			t.Errorf("Family = %v, want FamilyArray", b.Family)
		}
		if b.DBType != "ARRAY(Integer)" {
			t.Errorf("DBType = %q", b.DBType)
		}
		if !strings.Contains(b.APIType, "list[int]") {
			t.Errorf("APIType = %q", b.APIType)
		}
		if len(b.DBImports) != 1 {
			t.Errorf("expected one DBImport, got %v", b.DBImports)
		}
	})

	t.Run("SupportedNullable", func(t *testing.T) {
		b := bindArray(ir.BindOpts{ElementHead: "TEXT", NotNull: false})
		if b.DBType != "ARRAY(String)" {
			t.Errorf("DBType = %q", b.DBType)
		}
		// Nullable should differ from the NOT NULL form.
		nn := bindArray(ir.BindOpts{ElementHead: "TEXT", NotNull: true})
		if b.APIType == nn.APIType {
			t.Errorf("expected nullable APIType to differ: %q", b.APIType)
		}
	})

	t.Run("Unsupported", func(t *testing.T) {
		b := bindArray(ir.BindOpts{ElementHead: "UUID"})
		if b.Supported {
			t.Error("expected Supported false for UUID element")
		}
		if !strings.Contains(b.DBType, "unsupported array element: UUID") {
			t.Errorf("DBType = %q", b.DBType)
		}
	})
}

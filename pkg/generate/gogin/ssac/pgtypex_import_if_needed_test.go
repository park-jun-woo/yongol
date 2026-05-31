//ff:func feature=gen-gogin type=test control=sequence
//ff:what pgtypexImportIfNeeded 단위 테스트 (pgtypex bridge 컬럼이면 import 반환)
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestPgtypexImportIfNeeded(t *testing.T) {
	t.Run("nil column → nil", func(t *testing.T) {
		if got := pgtypexImportIfNeeded(nil); got != nil {
			t.Errorf("expected nil, got %v", got)
		}
	})
	t.Run("native column → nil", func(t *testing.T) {
		col := &ddl.Column{Name: "id", RawType: "BIGINT", NotNull: true}
		if got := pgtypexImportIfNeeded(col); got != nil {
			t.Errorf("native column should not need pgtypex, got %v", got)
		}
	})
	t.Run("pgtype bridge column → import", func(t *testing.T) {
		col := &ddl.Column{Name: "id", RawType: "UUID", NotNull: true}
		got := pgtypexImportIfNeeded(col)
		if len(got) != 1 || got[0] != `"github.com/park-jun-woo/ssac/pkg/pgtypex"` {
			t.Errorf("expected pgtypex import, got %v", got)
		}
	})
}

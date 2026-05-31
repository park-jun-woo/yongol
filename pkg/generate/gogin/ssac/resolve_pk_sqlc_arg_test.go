//ff:func feature=gen-gogin type=test control=sequence
//ff:what resolvePKSqlcArg 단위 테스트 (PK 바인딩으로 InsertExpr 래핑 + pgtypex import)
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestResolvePKSqlcArg(t *testing.T) {
	uuidCol := &ddl.Column{Name: "id", RawType: "UUID", NotNull: true}
	bigintCol := &ddl.Column{Name: "id", RawType: "BIGINT", NotNull: true}

	t.Run("nil column → unchanged", func(t *testing.T) {
		expr, imps := resolvePKSqlcArg(nil, "rid", false)
		if expr != "rid" || imps != nil {
			t.Errorf("got (%q,%v), want (rid,nil)", expr, imps)
		}
	})
	t.Run("alreadyPgtype → unchanged", func(t *testing.T) {
		expr, imps := resolvePKSqlcArg(uuidCol, "row.ID", true)
		if expr != "row.ID" || imps != nil {
			t.Errorf("got (%q,%v), want (row.ID,nil)", expr, imps)
		}
	})
	t.Run("native passthrough InsertExpr {var} → unchanged", func(t *testing.T) {
		expr, imps := resolvePKSqlcArg(bigintCol, "rid", false)
		if expr != "rid" || imps != nil {
			t.Errorf("got (%q,%v), want (rid,nil)", expr, imps)
		}
	})
	t.Run("uuid wraps with InsertExpr + pgtypex import", func(t *testing.T) {
		expr, imps := resolvePKSqlcArg(uuidCol, "rid", false)
		if expr != "pgtypex.ToPgUUID(rid)" {
			t.Errorf("expr = %q, want pgtypex.ToPgUUID(rid)", expr)
		}
		if len(imps) != 1 || imps[0] != `"github.com/park-jun-woo/ssac/pkg/pgtypex"` {
			t.Errorf("imports = %v, want single pgtypex", imps)
		}
	})
}

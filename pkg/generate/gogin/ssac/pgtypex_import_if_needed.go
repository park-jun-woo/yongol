//ff:func feature=gen-gogin type=util control=sequence
//ff:what pgtypexImportIfNeeded — DDL 컬럼이 pgtypex bridge 면 import 경로 반환

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/generate/gogin/types"
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// pgtypexImportIfNeeded returns the pgtypex import path slice when the
// column's binding requires it (NilCheckExpr non-empty). Returns nil
// otherwise so callers can append unconditionally.
func pgtypexImportIfNeeded(col *ddl.Column) []string {
	if col == nil {
		return nil
	}
	binding := types.MapPGType(*col)
	if binding.NilCheckExpr != "" {
		return []string{`"github.com/park-jun-woo/ssac/pkg/pgtypex"`}
	}
	return nil
}

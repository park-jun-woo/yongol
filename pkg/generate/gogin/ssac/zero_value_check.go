//ff:func feature=gen-gogin type=util control=sequence
//ff:what zeroValueCheckWithCol — 타입에 따른 zero/nil-value 비교 표현식 (pgtypex NilCheck 분기)

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/generate/gogin/types"
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// zeroValueCheckWithCol returns the nil/zero-value predicate for the
// given target expression using the DDL column's binding NilCheckExpr
// when available. Falls back to native zero comparison.
func zeroValueCheckWithCol(target string, col *ddl.Column) string {
	if col != nil {
		binding := types.MapPGType(*col)
		if binding.NilCheckExpr != "" {
			return types.Expand(binding.NilCheckExpr, "", "", target)
		}
	}
	return target + " == 0"
}

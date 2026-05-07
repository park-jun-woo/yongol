//ff:func feature=gen-gogin type=util control=sequence
//ff:what nonZeroCheckWithCol — 타입에 따른 non-zero predicate 표현식 (negated nil-check)

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/generate/gogin/types"
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// nonZeroCheckWithCol returns the non-zero predicate (negated nil-check).
func nonZeroCheckWithCol(target string, col *ddl.Column) string {
	if col != nil {
		binding := types.MapPGType(*col)
		if binding.NilCheckExpr != "" {
			return "!" + types.Expand(binding.NilCheckExpr, "", "", target)
		}
	}
	return target + " != 0"
}

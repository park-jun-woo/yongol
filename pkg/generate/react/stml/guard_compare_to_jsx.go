//ff:func feature=stml-gen type=util control=sequence dimension=1
//ff:what GuardCompare 노드(ref op value)를 JSX 비교식으로 변환한다
package stml

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// guardCompareToJSX renders a comparison node as a JSX expression. The reference
// path is emitted with optional chaining under dataVar (matching the legacy
// field=value output), and the value is single-quoted to match legacy strict
// equality output.
func guardCompareToJSX(expr *stml.GuardExpr, dataVar string) string {
	ref := fmt.Sprintf("%s.%s", dataVar, optionalChainPath(expr.Ref.Path()))
	return fmt.Sprintf("%s %s '%s'", ref, jsxCompareOp(expr.Op), expr.Value)
}

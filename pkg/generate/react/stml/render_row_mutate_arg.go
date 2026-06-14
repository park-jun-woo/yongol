//ff:func feature=stml-gen type=util control=iteration dimension=1
//ff:what 행 액션의 mutate() 호출 인자 객체 리터럴을 생성한다 (item.*/route.* 혼합, Number 래핑 포함)
package stml

import (
	"fmt"
	"strings"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// renderRowMutateArg builds the `{ key: value, ... }` argument object passed
// to mutation.mutate() at the row-action call site (inside the data-each map
// callback, where both useParams() names and `item` are in scope). Integer
// path parameters are wrapped with Number(...) like renderParamArgs, except
// item.<Field> sources whose field is already numeric in the response item
// schema (page-flow Phase006).
func renderRowMutateArg(a stmlparser.ActionBlock, pathParamTypes map[string]map[string]string, itemFieldTypes map[string]string) string {
	if len(a.Params) == 0 {
		return ""
	}
	var parts []string
	for _, p := range a.Params {
		expr := paramSourceExpr(p)
		if isIntegerParam(a.OperationID, p.Name, pathParamTypes) && !itemParamIsNumber(p, itemFieldTypes) {
			expr = wrapNumberArg(expr)
		}
		parts = append(parts, fmt.Sprintf("%s: %s", p.Name, expr))
	}
	return "{ " + strings.Join(parts, ", ") + " }"
}

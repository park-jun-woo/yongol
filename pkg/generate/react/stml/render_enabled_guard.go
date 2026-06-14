//ff:func feature=stml-gen type=generator control=iteration dimension=1
//ff:what optional route 파라미터에 의존하는 useQuery의 enabled 가드 문자열을 만든다 (없으면 빈 문자열, BUG-136)
package stml

import (
	"strings"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// renderEnabledGuard returns the "\n    enabled: <expr>," line for a useQuery
// whose fetch depends on an optional route param, or "" when every param is
// required (no guard, byte-identical to the pre-BUG-136 output). Integer params
// guard on Number.isFinite(Number(name)); other params on truthiness (!!name).
// Multiple optional params are AND-joined.
func renderEnabledGuard(f stmlparser.FetchBlock, pathParamTypes map[string]map[string]string) string {
	var guards []string
	for _, p := range f.Params {
		if !p.Optional {
			continue
		}
		v := paramSourceExpr(p)
		if isIntegerParam(f.OperationID, p.Name, pathParamTypes) {
			guards = append(guards, "Number.isFinite(Number("+v+"))")
		} else {
			guards = append(guards, "!!"+v)
		}
	}
	if len(guards) == 0 {
		return ""
	}
	return "\n    enabled: " + strings.Join(guards, " && ") + ","
}

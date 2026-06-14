//ff:func feature=stml-gen type=util control=iteration dimension=1
//ff:what optional route param을 소비하는 mutation 트리거를 막는 boolean 가드식을 만든다 (없으면 빈 문자열, BUG-136)
package stml

import (
	"strings"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// renderOptionalMutationGuard returns a JSX boolean expression that is true when
// any optional route param consumed by the action is absent, or "" when the
// action has no optional route param. An absent optional segment yields
// undefined from useParams(), so the guard tests "<name> == null" per param,
// OR-joined. Callers merge it into a Button `disabled` so an optional-param
// mutation never fires with NaN (BUG-136). The arg itself stays plain
// Number(...) for type fidelity (BUG-137) — this guard, not the arg, prevents
// the empty-value call. The Optional flag rides on ParamBind
// (markOptionalRouteParams), so no path-param type threading is needed.
func renderOptionalMutationGuard(a stmlparser.ActionBlock) string {
	var guards []string
	for _, p := range a.Params {
		if !p.Optional {
			continue
		}
		guards = append(guards, paramSourceExpr(p)+" == null")
	}
	if len(guards) == 0 {
		return ""
	}
	return strings.Join(guards, " || ")
}

//ff:func feature=validate type=helper control=selection dimension=1 topic=stml-statemachine
//ff:what collectEqualComparePairs — 가드 AST에서 무조건 요구되는 "=" 비교쌍만 수집 (!·||·비등호는 침묵)

package stml_statemachine

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// collectEqualComparePairs walks the guard AST and returns the (model,
// value) pairs of "=" comparisons that the guard requires unconditionally:
// "&&" branches recurse into both sides, while "||" alternatives, negations
// ("!"), and non-equality operators yield nothing — they do not pin the
// state to a single value, so TM-23 treats them as not comparable and stays
// silent. A nil node yields nil.
func collectEqualComparePairs(expr *stml.GuardExpr) []comparePair {
	if expr == nil {
		return nil
	}
	switch expr.Kind {
	case stml.GuardBinary:
		if expr.Op != "&&" {
			return nil
		}
		return append(collectEqualComparePairs(expr.Left), collectEqualComparePairs(expr.Right)...)
	case stml.GuardGroup:
		return collectEqualComparePairs(expr.Operand)
	case stml.GuardCompare:
		if expr.Op != "=" {
			return nil
		}
		return []comparePair{{Model: expr.Ref.Model, Value: expr.Value}}
	default:
		return nil
	}
}

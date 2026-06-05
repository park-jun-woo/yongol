//ff:func feature=validate type=helper control=selection dimension=1 topic=stml-statemachine
//ff:what collectComparePairs — 가드 AST를 순회해 모든 GuardCompare의 (model, value) 쌍을 DOM 순서로 수집

package stml_statemachine

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// collectComparePairs walks the guard AST and returns the (model, value) pair of
// every GuardCompare leaf in left-to-right order. Binary, unary, and group nodes
// recurse; lifecycle leaves contribute nothing (no compared value). A nil node
// yields nil.
func collectComparePairs(expr *stml.GuardExpr) []comparePair {
	if expr == nil {
		return nil
	}
	switch expr.Kind {
	case stml.GuardBinary:
		return append(collectComparePairs(expr.Left), collectComparePairs(expr.Right)...)
	case stml.GuardUnary, stml.GuardGroup:
		return collectComparePairs(expr.Operand)
	case stml.GuardCompare:
		return []comparePair{{Model: expr.Ref.Model, Value: expr.Value}}
	default:
		return nil
	}
}

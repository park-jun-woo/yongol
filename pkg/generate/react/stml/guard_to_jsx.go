//ff:func feature=stml-gen type=util control=selection dimension=1
//ff:what GuardExpr AST를 JSX 조건 표현식 문자열로 재귀 변환한다
package stml

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// guardToJSX converts a guard AST node into a JSX boolean expression string,
// recursing into combining, unary, and group nodes and delegating leaf
// conversion (comparison / lifecycle) to helpers.
func guardToJSX(expr *stml.GuardExpr, dataVar string) string {
	switch expr.Kind {
	case stml.GuardBinary:
		left := guardToJSX(expr.Left, dataVar)
		right := guardToJSX(expr.Right, dataVar)
		return left + " " + expr.Op + " " + right
	case stml.GuardUnary:
		return "!" + guardToJSX(expr.Operand, dataVar)
	case stml.GuardGroup:
		return "(" + guardToJSX(expr.Operand, dataVar) + ")"
	case stml.GuardLifecycle:
		return guardLifecycleToJSX(expr, dataVar)
	default:
		return guardCompareToJSX(expr, dataVar)
	}
}

//ff:func feature=stml-parse type=util control=selection dimension=1
//ff:what GuardExpr AST를 순회해 모든 GuardRef(model.Field)를 DOM 순서로 수집한다
package stml

// CollectRefs walks the guard AST and returns every GuardRef (model.Field) it
// contains, in left-to-right order. Compare and lifecycle leaves contribute
// their Ref; binary, unary, and group nodes recurse. A nil receiver yields nil.
func (expr *GuardExpr) CollectRefs() []GuardRef {
	if expr == nil {
		return nil
	}
	switch expr.Kind {
	case GuardBinary:
		return append(expr.Left.CollectRefs(), expr.Right.CollectRefs()...)
	case GuardUnary, GuardGroup:
		return expr.Operand.CollectRefs()
	default:
		return []GuardRef{expr.Ref}
	}
}

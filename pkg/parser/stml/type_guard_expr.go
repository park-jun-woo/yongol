//ff:type feature=stml-parse type=model
//ff:what GuardExpr — 가드 문법(§3.4 EBNF) AST 노드 (Kind에 따라 활성 필드가 달라짐)
package stml

// GuardExpr is a node in the guard expression AST. The active fields depend on
// Kind:
//
//	GuardBinary    → Op ("&&"|"||"), Left, Right
//	GuardUnary     → Operand
//	GuardCompare   → Ref, Op (= != > < >= <=), Value
//	GuardLifecycle → Ref, Lifecycle (loading|error|empty)
//	GuardGroup     → Operand
type GuardExpr struct {
	Kind      GuardKind
	Op        string     // logical (&&,||) or comparison (=,!=,>,<,>=,<=) operator
	Left      *GuardExpr // GuardBinary left operand
	Right     *GuardExpr // GuardBinary right operand
	Operand   *GuardExpr // GuardUnary / GuardGroup inner expression
	Ref       GuardRef   // GuardCompare / GuardLifecycle reference
	Value     string     // GuardCompare right-hand literal value
	Lifecycle string     // GuardLifecycle keyword (loading|error|empty)
}

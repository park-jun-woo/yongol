//ff:type feature=stml-parse type=model
//ff:what GuardKind — 가드 AST 노드 종류 enum (이항/단항/비교/생명주기/그룹)
package stml

// GuardKind discriminates the variant of a GuardExpr node.
type GuardKind int

const (
	// GuardBinary is a logical combination: Left Op Right where Op is "&&" or "||".
	GuardBinary GuardKind = iota
	// GuardUnary is a negation: "!" Operand.
	GuardUnary
	// GuardCompare is a comparison: Ref Op Value (op in = != > < >= <=).
	GuardCompare
	// GuardLifecycle is a fetch lifecycle test: Ref "." Lifecycle (loading|error|empty).
	GuardLifecycle
	// GuardGroup is a parenthesized sub-expression: "(" guard ")".
	GuardGroup
)

//ff:func feature=validate type=util control=sequence topic=ssac-structural
//ff:what isBoolPredicateSignature — FuncSpec returns exactly one bool

package ssac

import "github.com/park-jun-woo/yongol/pkg/parser/funcspec"

// isBoolPredicateSignature reports whether spec declares a single `bool`
// return value, the canonical shape required for @eval predicate guards.
func isBoolPredicateSignature(spec *funcspec.FuncSpec) bool {
	return len(spec.ReturnTypes) == 1 && spec.ReturnTypes[0] == "bool"
}

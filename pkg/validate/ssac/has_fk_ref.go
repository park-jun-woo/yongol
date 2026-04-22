//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-structural
//ff:what hasFKRef — @get args/inputs 중 declared 변수 참조(다른 모델)인지 판정

package ssac

import (
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// hasFKRef reports whether seq references a previously declared variable
// whose type differs from the current @get model (i.e., a foreign-key
// reference). Implicit sources (request/currentUser/query/message) are
// skipped.
func hasFKRef(seq parsessac.Sequence, declared map[string]bool, types map[string]string, getModel string) bool {
	for _, arg := range seq.Args {
		ref := arg.Source
		if ref == "" || isImplicitVar(ref) {
			continue
		}
		if declared[ref] && types[ref] != getModel {
			return true
		}
	}
	for _, val := range seq.Inputs {
		ref := inputValueRefBase(val)
		if ref == "" || isImplicitVar(ref) {
			continue
		}
		if declared[ref] && types[ref] != getModel {
			return true
		}
	}
	return false
}

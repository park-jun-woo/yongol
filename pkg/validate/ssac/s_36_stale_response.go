//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-36 — @put/@delete 가 변경한 변수를 re-@get 없이 @response 에 사용

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s36StaleResponse validates S-36: when @put/@delete mutates the DB row
// backing a variable, any subsequent @response that references the variable
// without a re-@get assignment is stale (the in-memory struct is out of date).
//
// Algorithm (sequence-order aware):
//  1. Track variable → type via Result.Var assignments
//  2. On @put/@delete of model M, mark every variable whose type is M stale
//  3. On @get re-assignment of var V, clear stale[V]
//  4. @response referencing a stale var → WARNING
func s36StaleResponse(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		varType := map[string]string{}
		stale := map[string]bool{}
		for i, seq := range fn.Sequences {
			s36TrackAssignment(seq, varType, stale)
			s36MarkStaleAfterMutation(seq, varType, stale)
			diags = append(diags, s36CheckResponseStale(fn, i, seq, stale)...)
		}
	}
	return diags
}

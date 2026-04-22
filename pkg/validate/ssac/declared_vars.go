//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-structural
//ff:what declaredVars — seqIdx 이전까지 선언된 result 변수 집합 (+ subscribe message)

package ssac

import (
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// declaredVars walks the sequence list and returns the set of result variable
// names declared up to (but not including) seqIdx, plus the sub message var
// when fn is a subscribe function.
func declaredVars(fn parsessac.ServiceFunc, upto int) map[string]bool {
	out := make(map[string]bool)
	if fn.Subscribe != nil {
		out["message"] = true
	}
	for i := 0; i < upto && i < len(fn.Sequences); i++ {
		seq := fn.Sequences[i]
		if seq.Result != nil && seq.Result.Var != "" {
			out[seq.Result.Var] = true
		}
	}
	return out
}

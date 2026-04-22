//ff:func feature=validate type=util control=sequence topic=ssac-structural
//ff:what s36TrackAssignment — Result.Var 기록 + stale 플래그 초기화

package ssac

import parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"

// s36TrackAssignment records the variable → type binding for the current
// sequence (Result.Var ← Result.Type) and resets its stale flag because the
// variable was freshly assigned.
func s36TrackAssignment(seq parsessac.Sequence, varType map[string]string, stale map[string]bool) {
	if seq.Result == nil || seq.Result.Var == "" || seq.Result.Type == "" {
		return
	}
	t := stripTypePrefix(seq.Result.Type)
	varType[seq.Result.Var] = t
	stale[seq.Result.Var] = false
}

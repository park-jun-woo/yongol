//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-structural
//ff:what s36MarkStaleAfterMutation — @put/@delete 대상 model 변수 stale 표시

package ssac

import parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"

// s36MarkStaleAfterMutation flags every tracked variable whose backing model
// matches the sequence's @put/@delete target as stale.
func s36MarkStaleAfterMutation(seq parsessac.Sequence, varType map[string]string, stale map[string]bool) {
	if seq.Type != "put" && seq.Type != "delete" {
		return
	}
	model := extractModel(seq)
	if model == "" {
		return
	}
	for v, t := range varType {
		if t == model {
			stale[v] = true
		}
	}
}

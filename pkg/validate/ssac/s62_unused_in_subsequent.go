//ff:func feature=validate type=util control=iteration dimension=2 topic=ssac-structural
//ff:what s62unusedInSubsequent — varName 이 start 이후 sequences 에서 참조되지 않으면 true

package ssac

import (
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// s62unusedInSubsequent returns true when varName does not appear in any of
// the sequences starting at index start (Inputs values, Fields values, Target).
func s62unusedInSubsequent(varName string, seqs []parsessac.Sequence, start int) bool {
	for _, seq := range seqs[start:] {
		for _, v := range seq.Inputs {
			if s62prefix(v) == varName {
				return false
			}
		}
		for _, v := range seq.Fields {
			if s62prefix(v) == varName {
				return false
			}
		}
		if seq.Target != "" && s62prefix(seq.Target) == varName {
			return false
		}
	}
	return true
}

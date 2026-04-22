//ff:func feature=validate type=util control=iteration dimension=1 topic=states
//ff:what SSaC sequence 리스트에 @state 시퀀스가 있는지 확인

package ssac_statemachine

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// hasStateSequence checks whether any sequence is a @state sequence.
func hasStateSequence(seqs []ssac.Sequence) bool {
	for _, seq := range seqs {
		if seq.Type == "state" {
			return true
		}
	}
	return false
}

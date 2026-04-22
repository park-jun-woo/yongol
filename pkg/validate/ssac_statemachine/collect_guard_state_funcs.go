//ff:func feature=validate type=util control=iteration dimension=1 topic=states
//ff:what @state 시퀀스가 있는 함수명 수집

package ssac_statemachine

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// collectGuardStateFuncs collects function names that have @state sequences.
func collectGuardStateFuncs(funcs []ssac.ServiceFunc) map[string]bool {
	result := make(map[string]bool, len(funcs))
	for _, fn := range funcs {
		if hasStateSequence(fn.Sequences) {
			result[fn.Name] = true
		}
	}
	return result
}

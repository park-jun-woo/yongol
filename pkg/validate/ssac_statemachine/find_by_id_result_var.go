//ff:func feature=validate type=util control=iteration dimension=1 topic=states
//ff:what findByIDResultVar — `@get <Model>.FindByID(...)` 시퀀스의 결과 변수명 탐색

package ssac_statemachine

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// findByIDResultVar returns the bound variable name of the `@get
// <Model>.FindByID(...)` sequence and reports whether such a sequence was
// found. When the `@get` has no result variable the empty string is returned
// alongside ok=true, which still signals "resource is read" for detection
// purposes; callers pick a fallback name for diagnostics.
func findByIDResultVar(seqs []ssac.Sequence, model string) (string, bool) {
	if model == "" {
		return "", false
	}
	wanted := model + ".FindByID"
	for _, seq := range seqs {
		if seq.Type != "get" {
			continue
		}
		if seq.Model != wanted {
			continue
		}
		if seq.Result != nil {
			return seq.Result.Var, true
		}
		return "", true
	}
	return "", false
}

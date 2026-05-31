//ff:func feature=gen-ir type=test control=sequence
//ff:what convertDelete/convertPut/convertEmpty/convertExists/matchFollowingGuard/resolveVar/convertInputsToFieldArgs
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestConvertExists(t *testing.T) {
	op := convertExists(ssac.Sequence{Target: "dup", Message: "conflict"})
	if op.Kind != OpExists || op.Exists.VarName != "dup" || op.Exists.StatusCode != 409 {
		t.Errorf("op = %+v", op.Exists)
	}
	withStatus := convertExists(ssac.Sequence{Target: "y", ErrStatus: 422})
	if withStatus.Exists.StatusCode != 422 {
		t.Errorf("status = %d, want 422", withStatus.Exists.StatusCode)
	}
}

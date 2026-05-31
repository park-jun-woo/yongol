//ff:func feature=gen-ir type=test control=sequence
//ff:what convertDelete/convertPut/convertEmpty/convertExists/matchFollowingGuard/resolveVar/convertInputsToFieldArgs
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestConvertEmpty(t *testing.T) {
	// explicit status
	op := convertEmpty(ssac.Sequence{Target: "course", Message: "not found", ErrStatus: 410})
	if op.Kind != OpEmpty || op.Empty.VarName != "course" || op.Empty.Message != "not found" || op.Empty.StatusCode != 410 {
		t.Errorf("op = %+v", op.Empty)
	}
	// default status 404
	def := convertEmpty(ssac.Sequence{Target: "x"})
	if def.Empty.StatusCode != 404 {
		t.Errorf("default status = %d, want 404", def.Empty.StatusCode)
	}
}

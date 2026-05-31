//ff:func feature=gen-ir type=test control=sequence
//ff:what convertDelete/convertPut/convertEmpty/convertExists/matchFollowingGuard/resolveVar/convertInputsToFieldArgs
package ir

import (
	"testing"
)

func TestConvertInputsToFieldArgs(t *testing.T) {
	if got := convertInputsToFieldArgs(nil); got != nil {
		t.Errorf("nil inputs = %+v, want nil", got)
	}
	args := convertInputsToFieldArgs(map[string]string{"Zeta": "request.Z", "Alpha": "request.A"})
	if len(args) != 2 {
		t.Fatalf("len = %d, want 2", len(args))
	}
	// keys sorted deterministically: Alpha before Zeta
	if args[0].Key != "Alpha" || args[1].Key != "Zeta" {
		t.Errorf("keys not sorted: %q, %q", args[0].Key, args[1].Key)
	}
}

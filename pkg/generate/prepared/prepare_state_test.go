//ff:func feature=generate type=test control=sequence
//ff:what TestPrepareState_BUG008 — manifest·SSaC 모두 session을 안 쓰는 경우 panic 없이 nil

package prepared

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestPrepareState_BUG008 pins the BUG-008 regression at the prepare
// layer: with no manifest.session and no SSaC session.* calls the
// resulting State must leave ActiveBackends.Session nil so no session
// block reaches codegen.
func TestPrepareState_BUG008(t *testing.T) {
	fs := &yongol.Fullstack{}
	p := New(fs)
	if p.ActiveBackends.Session != nil {
		t.Fatalf("ActiveBackends.Session=%+v; expected nil for empty Fullstack", p.ActiveBackends.Session)
	}
	if p.ActiveBackends.Cache != nil || p.ActiveBackends.File != nil || p.ActiveBackends.Queue != nil {
		t.Fatalf("ActiveBackends should be all nil, got %+v", p.ActiveBackends)
	}
	if p.Auth.Present {
		t.Fatalf("Auth.Present=true for empty Fullstack; expected false")
	}
}

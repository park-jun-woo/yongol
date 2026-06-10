//ff:func feature=generate type=test control=sequence
//ff:what TestPrepareAuthCsrfRequired_Absent — manifest.auth 미선언 → CsrfRequired=false

package prepared

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestPrepareAuthCsrfRequired_Absent asserts that a manifest with no
// backend.auth subtree derives Auth{Present:false, CsrfRequired:false}.
// Projects that do not declare auth at all must not receive CSRF
// wiring — there is no session to protect.
func TestPrepareAuthCsrfRequired_Absent(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{Auth: nil},
		},
	}
	got := AuthFor(fs)
	if got.Present {
		t.Fatalf("Present=true; expected false when auth is absent")
	}
	if got.CsrfRequired {
		t.Fatalf("CsrfRequired=true; expected false when auth is absent")
	}
}

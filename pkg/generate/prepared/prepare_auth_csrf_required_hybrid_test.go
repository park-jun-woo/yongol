//ff:func feature=generate type=test control=sequence
//ff:what TestPrepareAuthCsrfRequired_Hybrid — manifest.auth.mode=hybrid → CsrfRequired=true

package prepared

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestPrepareAuthCsrfRequired_Hybrid asserts that hybrid auth mode
// also derives CsrfRequired=true. Bearer-header requests still
// bypass CSRF inside the middleware via HybridBearerSkip, but the
// block itself must be emitted so the cookie-session fork stays
// protected.
func TestPrepareAuthCsrfRequired_Hybrid(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Auth: &pmanifest.Auth{Mode: "hybrid"},
			},
		},
	}
	got := AuthFor(fs)
	if !got.Present {
		t.Fatalf("Present=false; expected true")
	}
	if got.Mode != "hybrid" {
		t.Fatalf("Mode=%q; expected %q", got.Mode, "hybrid")
	}
	if !got.CsrfRequired {
		t.Fatalf("CsrfRequired=false; expected true for hybrid mode")
	}
}

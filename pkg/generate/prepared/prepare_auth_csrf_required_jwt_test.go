//ff:func feature=generate type=test control=sequence
//ff:what TestPrepareAuthCsrfRequired_Jwt — manifest.auth.type=jwt → CsrfRequired=false

package prepared

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestPrepareAuthCsrfRequired_Jwt pins BUG-013: a manifest with
// auth.type=jwt and no explicit mode must resolve to bearer with
// CsrfRequired=false — JWT-only projects are CSRF-immune and must not
// emit the CSRF middleware.
func TestPrepareAuthCsrfRequired_Jwt(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Auth: &pmanifest.Auth{Type: "jwt"},
			},
		},
	}
	got := authFor(fs)
	if !got.Present {
		t.Fatalf("Present=false; expected true")
	}
	if got.Mode != "bearer" {
		t.Fatalf("Mode=%q; expected %q", got.Mode, "bearer")
	}
	if got.CsrfRequired {
		t.Fatalf("CsrfRequired=true; expected false for bearer mode")
	}
}

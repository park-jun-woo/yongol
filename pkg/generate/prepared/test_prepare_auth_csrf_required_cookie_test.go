//ff:func feature=generate type=test control=sequence
//ff:what TestPrepareAuthCsrfRequired_Cookie — manifest.auth.mode=cookie → CsrfRequired=true

package prepared

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestPrepareAuthCsrfRequired_Cookie asserts that cookie auth mode
// derives CsrfRequired=true — the classic double-submit defence is
// mandatory for cookie-session projects.
func TestPrepareAuthCsrfRequired_Cookie(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Auth: &pmanifest.Auth{Mode: "cookie"},
			},
		},
	}
	got := authFor(fs)
	if !got.Present {
		t.Fatalf("Present=false; expected true")
	}
	if got.Mode != "cookie" {
		t.Fatalf("Mode=%q; expected %q", got.Mode, "cookie")
	}
	if !got.CsrfRequired {
		t.Fatalf("CsrfRequired=false; expected true for cookie mode")
	}
}

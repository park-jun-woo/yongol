//ff:func feature=manifest type=test control=sequence
//ff:what TestManifestHelpers — unit tests for the pure manifest parser helper functions
package manifest

import (
	"testing"
)

func TestResolvedMode(t *testing.T) {
	if got := (&Auth{}).ResolvedMode(); got != "cookie" {
		t.Errorf("empty mode → %q, want cookie", got)
	}
	if got := (&Auth{Mode: "bearer"}).ResolvedMode(); got != "bearer" {
		t.Errorf("explicit bearer → %q", got)
	}
	var nilAuth *Auth
	if got := nilAuth.ResolvedMode(); got != "cookie" {
		t.Errorf("nil auth → %q, want cookie", got)
	}
}

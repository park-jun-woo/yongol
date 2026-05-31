//ff:func feature=validate type=test control=sequence topic=config-check
//ff:what TestSSaCManifestHelpers — unit tests for the pure ssac_manifest helper functions
package ssac_manifest

import (
	"testing"
)

func TestSortedClaimFields(t *testing.T) {
	got := sortedClaimFields(map[string]bool{"role": true, "id": true, "email": true})
	if got != "email, id, role" {
		t.Errorf("got %q, want 'email, id, role'", got)
	}
	if got := sortedClaimFields(map[string]bool{}); got != "" {
		t.Errorf("empty → %q", got)
	}
}

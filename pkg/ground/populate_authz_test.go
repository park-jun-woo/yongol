//ff:func feature=rule type=test control=iteration dimension=1
//ff:what populateAuthz — default CheckRequest 필드 집합 검증

package ground

import (
	"testing"
)

// TestPopulateAuthz_Defaults verifies the default field set when no custom
// authz package is configured.
func TestPopulateAuthz_Defaults(t *testing.T) {
	fs := newMinimalFullstack()
	g := newGround()

	populateAuthz(g, fs)

	set := g.Lookup["Authz.checkRequest"]
	for _, f := range []string{"Action", "Resource", "UserID", "Role", "ResourceID"} {
		if !set[f] {
			t.Errorf("Authz.checkRequest missing %q: %v", f, set)
		}
	}
}

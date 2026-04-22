//ff:func feature=rule type=test control=sequence dimension=1
//ff:what Build — 최소 Fullstack에서 모든 populate가 panic 없이 실행되고 기본 키를 채운다

package ground

import (
	"testing"
)

// TestBuild_FromMinimalFullstack ensures Build succeeds on a fully empty
// Fullstack (all SSOTs absent). populate_* must be nil-safe; core keys such as
// "go.reserved" should still be populated.
func TestBuild_FromMinimalFullstack(t *testing.T) {
	fs := newMinimalFullstack()
	g := Build(fs)
	if g == nil {
		t.Fatal("Build returned nil")
	}

	// populateGoReservedWords is unconditional — its key must be present.
	if !g.Lookup["go.reserved"]["func"] {
		t.Errorf("go.reserved missing 'func' — populateGoReservedWords regression: %v",
			g.Lookup["go.reserved"])
	}
	// populateAuthz defaults are always emitted when authz package is unset.
	if !g.Lookup["Authz.checkRequest"]["Action"] {
		t.Errorf("Authz.checkRequest defaults missing: %v", g.Lookup["Authz.checkRequest"])
	}
}

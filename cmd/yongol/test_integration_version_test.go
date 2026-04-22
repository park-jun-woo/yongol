//ff:func feature=cli type=test control=sequence
//ff:what test: version 서브커맨드 end-to-end 1 케이스 (basic)

package main

import (
	"strings"
	"testing"
)

// TestIntegrationVersion_Basic runs `yongol version` and expects exit 0
// with stdout containing the literal `yongol` plus the build-time
// Version string (defaults to `dev` outside a release build).
func TestIntegrationVersion_Basic(t *testing.T) {
	stdout, _, err := runCmd(t, "version")
	if err != nil {
		t.Fatalf("version: unexpected error: %v", err)
	}
	if !strings.Contains(stdout, "yongol") {
		t.Errorf("expected stdout to contain `yongol`, got: %q", stdout)
	}
	if !strings.Contains(stdout, Version) {
		t.Errorf("expected stdout to contain Version=%q, got: %q", Version, stdout)
	}
}

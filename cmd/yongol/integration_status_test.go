//ff:func feature=cli type=test control=sequence
//ff:what status happy — zenflow specs 로 status dashboard 성공 시나리오

package main

import (
	"strings"
	"testing"
)

// TestIntegrationStatus_Happy runs `yongol status <zenflow-specs>` and
// expects exit 0 with an SSOT Summary header. status is a read-only
// dashboard, so no arts-dir means no drift / artifacts sections.
func TestIntegrationStatus_Happy(t *testing.T) {
	specs := zenflowSpecsDir(t)
	stdout, _, err := runCmd(t, "status", specs)
	if err != nil {
		t.Fatalf("status happy: unexpected error: %v\nstdout:\n%s", err, stdout)
	}
	if !strings.Contains(stdout, "SSOT Summary") {
		t.Errorf("expected stdout to contain `SSOT Summary`, got:\n%s", stdout)
	}
	// OpenAPI endpoints row is always rendered for a populated zenflow tree.
	if !strings.Contains(stdout, "OpenAPI") {
		t.Errorf("expected stdout to mention OpenAPI in summary, got:\n%s", stdout)
	}
}

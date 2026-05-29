//ff:func feature=cli type=test control=sequence
//ff:what chain happy — ExecuteWorkflow 추적 성공 및 header/링크 검증

package main

import (
	"strings"
	"testing"
)

// TestIntegrationChain_Happy traces the ExecuteWorkflow operationId end-to-end
// through zenflow specs. Expects exit 0 and a formatted header —
// `── Feature Chain: ExecuteWorkflow ──` — plus at least the OpenAPI link.
func TestIntegrationChain_Happy(t *testing.T) {
	specs := zenflowSpecsDir(t)
	stdout, _, err := runCmd(t, "chain", "ExecuteWorkflow", specs)
	if err != nil {
		t.Fatalf("chain happy: unexpected error: %v\nstdout:\n%s", err, stdout)
	}
	if !strings.Contains(stdout, "Feature Chain:") {
		t.Errorf("expected stdout to contain `Feature Chain:`, got:\n%s", stdout)
	}
	if !strings.Contains(stdout, "ExecuteWorkflow") {
		t.Errorf("expected stdout to mention ExecuteWorkflow, got:\n%s", stdout)
	}
	if !strings.Contains(stdout, "OpenAPI") {
		t.Errorf("expected stdout to list OpenAPI link, got:\n%s", stdout)
	}
}

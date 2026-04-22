//ff:func feature=cli type=test control=iteration dimension=1
//ff:what test: chain 서브커맨드 end-to-end 2 케이스 (happy / unknown-op)

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

// TestIntegrationChain_UnknownOp passes a non-existent operationId and
// expects exit 1 with a `not found in OpenAPI` error. Chain.Chain fails
// fast when traceOpenAPI yields nil.
func TestIntegrationChain_UnknownOp(t *testing.T) {
	specs := zenflowSpecsDir(t)
	_, _, err := runCmd(t, "chain", "DefinitelyNoSuchOp", specs)
	if err == nil {
		t.Fatal("expected error for unknown operationId, got nil")
	}
	if isUsageError(err) {
		t.Fatalf("unknown-op should be exit 1, not usage error: %v", err)
	}
	if !strings.Contains(err.Error(), "not found in OpenAPI") {
		t.Errorf("expected error to mention `not found in OpenAPI`, got: %v", err)
	}
}

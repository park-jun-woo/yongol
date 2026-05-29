//ff:func feature=cli type=test control=sequence
//ff:what chain unknown-op — 존재하지 않는 operationId 시 exit 1 + `not found in OpenAPI`

package main

import (
	"strings"
	"testing"
)

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

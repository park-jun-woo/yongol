//ff:func feature=chain type=test control=iteration dimension=1
//ff:what Chain end-to-end 를 zenflow dummy specs 로 검증 (OpenAPI/SSaC/DDL 최소 포함)
package chain

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestChain_Zenflow_ExecuteWorkflow exercises Chain end-to-end on the zenflow
// dummy specs. Skipped when the dummy directory is not present.
func TestChain_Zenflow_ExecuteWorkflow(t *testing.T) {
	specsDir := zenflowSpecsDir(t)
	if specsDir == "" {
		t.Skip("zenflow dummy specs not available")
	}
	detected, err := yongol.DetectSSOTs(specsDir)
	if err != nil {
		t.Fatalf("DetectSSOTs: %v", err)
	}
	fs := yongol.ParseAll(specsDir, detected)
	if fs.OpenAPIDoc == nil {
		t.Skip("zenflow OpenAPI not parsed")
	}

	links, err := Chain(fs, "ExecuteWorkflow")
	if err != nil {
		t.Fatalf("Chain error: %v", err)
	}
	if len(links) == 0 {
		t.Fatal("expected at least one link")
	}
	kinds := map[string]bool{}
	for _, l := range links {
		kinds[l.Kind] = true
	}
	for _, want := range []string{"OpenAPI", "SSaC", "DDL"} {
		if !kinds[want] {
			t.Errorf("missing expected kind %q (kinds seen: %v)", want, kinds)
		}
	}
}

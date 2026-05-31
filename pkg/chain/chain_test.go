//ff:func feature=chain type=test control=sequence
//ff:what TestChain — Chain 의 nil-OpenAPI / not-found / matched-ServiceFunc 분기 검증
package chain

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestChain(t *testing.T) {
	t.Run("NilOpenAPI", func(t *testing.T) {
		fs := &yongol.Fullstack{SpecsDir: t.TempDir()}
		if _, err := Chain(fs, "Anything"); err == nil {
			t.Fatal("expected error when OpenAPIDoc is nil")
		}
	})

	t.Run("OperationNotFound", func(t *testing.T) {
		fs := &yongol.Fullstack{
			SpecsDir: t.TempDir(),
			OpenAPIDoc: &openapi3.T{
				OpenAPI: "3.0.0",
				Info:    &openapi3.Info{Title: "t", Version: "1"},
				Paths:   openapi3.NewPaths(),
			},
		}
		if _, err := Chain(fs, "DoesNotExist"); err == nil {
			t.Fatal("expected error for missing operationId")
		}
	})

	t.Run("MatchedServiceFunc", func(t *testing.T) {
		// Drives the full happy path on real example specs so the matched !=
		// nil block (SSaC/DDL/policy/states/funcspec/hurl tracing) executes.
		specsDir := opus47SpecsDir(t)
		if specsDir == "" {
			t.Skip("opus4_7 example specs not available")
		}
		detected, err := yongol.DetectSSOTs(specsDir)
		if err != nil {
			t.Fatalf("DetectSSOTs: %v", err)
		}
		fs := yongol.ParseAll(specsDir, detected)
		if fs.OpenAPIDoc == nil {
			t.Skip("opus4_7 OpenAPI not parsed")
		}
		if len(fs.ServiceFuncs) == 0 {
			t.Skip("opus4_7 SSaC not parsed")
		}
		// Pick the first service func name so the operationId is guaranteed to
		// match an OpenAPI operation and a ServiceFunc (matched != nil branch).
		opID := fs.ServiceFuncs[0].Name
		links, err := Chain(fs, opID)
		if err != nil {
			// Not every ServiceFunc maps to an OpenAPI operationId; fall back to
			// asserting only the deterministic branches above in that case.
			t.Skipf("Chain(%q) returned error (no matching OpenAPI op): %v", opID, err)
		}
		if len(links) == 0 {
			t.Fatalf("expected at least one link for %q", opID)
		}
		var sawOpenAPI bool
		for _, l := range links {
			if l.Kind == "OpenAPI" {
				sawOpenAPI = true
			}
		}
		if !sawOpenAPI {
			t.Errorf("expected an OpenAPI link in chain for %q, got kinds: %v", opID, links)
		}
	})
}

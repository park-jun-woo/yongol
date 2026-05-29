//ff:func feature=chain type=test control=iteration dimension=2
//ff:what traceOpenAPI 가 operationId 매칭 시 Link / Paths nil 또는 미매칭 시 nil 을 반환하는지 검증
package chain

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestTraceOpenAPI(t *testing.T) {
	specsDir := t.TempDir()
	apiDir := filepath.Join(specsDir, "api")
	if err := os.MkdirAll(apiDir, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	yaml := "paths:\n  /courses/{id}:\n    get:\n      operationId: GetCourse\n"
	if err := os.WriteFile(filepath.Join(apiDir, "openapi.yaml"), []byte(yaml), 0o644); err != nil {
		t.Fatalf("write: %v", err)
	}

	paths := openapi3.NewPaths()
	paths.Set("/courses/{id}", &openapi3.PathItem{
		Get: &openapi3.Operation{OperationID: "GetCourse"},
	})
	doc := &openapi3.T{Paths: paths}

	link := traceOpenAPI(doc, "GetCourse", specsDir)
	if link == nil {
		t.Fatal("expected non-nil link")
	}
	if link.Kind != "OpenAPI" || link.File != "api/openapi.yaml" {
		t.Errorf("link fields: %+v", link)
	}
	if link.Summary != "GET /courses/{id}" {
		t.Errorf("summary: %q", link.Summary)
	}
	if link.Line != 4 {
		t.Errorf("line: got %d, want 4", link.Line)
	}

	// Unknown operationId → nil.
	if traceOpenAPI(doc, "Nope", specsDir) != nil {
		t.Error("expected nil for unknown operationId")
	}

	// Paths nil → nil.
	if traceOpenAPI(&openapi3.T{}, "GetCourse", specsDir) != nil {
		t.Error("expected nil when Paths is nil")
	}
}

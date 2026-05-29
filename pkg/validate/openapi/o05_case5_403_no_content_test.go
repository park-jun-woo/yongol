//ff:func feature=validate type=test control=sequence topic=response-body-required
//ff:what O-5 — 403 (description 만) 은 ERROR 1개 + 메시지·라인 검증

package openapi

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestO05_Case5_403NoContent(t *testing.T) {
	op := opWithResponses("Op5", map[string]*openapi3.ResponseRef{
		"403": emptyResponse("Forbidden"),
	})
	paths := openapi3.NewPaths(openapi3.WithPath("/x", &openapi3.PathItem{Get: op}))
	fs := newFullstackWithPaths(paths)
	fs.OpenAPILines.Operations = map[string]int{"Op5": 17}

	diags := o05ResponseBodyRequired(fs)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d: %+v", len(diags), diags)
	}
	got := diags[0]
	if !strings.Contains(got.Message, "[O-5]") {
		t.Errorf("message missing rule id: %q", got.Message)
	}
	if !strings.Contains(got.Message, "403") {
		t.Errorf("message missing status: %q", got.Message)
	}
	if !strings.Contains(got.Message, "Op5") {
		t.Errorf("message missing operationId: %q", got.Message)
	}
	if got.File != "api/openapi.yaml" {
		t.Errorf("file = %q, want api/openapi.yaml", got.File)
	}
	if got.Line != 17 {
		t.Errorf("line = %d, want 17 (OperationLine)", got.Line)
	}
}

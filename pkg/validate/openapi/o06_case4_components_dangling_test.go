//ff:func feature=validate type=test control=sequence topic=openapi-structural
//ff:what O-6 — components 스키마 dangling required → ERROR + SchemaLine 폴백

package openapi

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestO06_Case4_ComponentsDangling(t *testing.T) {
	doc := &openapi3.T{Components: &openapi3.Components{Schemas: openapi3.Schemas{
		"Workflow": o06SchemaWithRequired([]string{"id"}, []string{"phantom"}),
	}}}
	fs := &yongol.Fullstack{
		OpenAPIDoc:   doc,
		OpenAPILines: &oapiparser.LineIndex{Schemas: map[string]int{"Workflow": 42}, Paths: map[string]int{}, Operations: map[string]int{}},
	}
	diags := o06RequiredPropertyConsistency(fs)
	if len(diags) != 1 {
		t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
	}
	if !strings.Contains(diags[0].Message, "O-6") || !strings.Contains(diags[0].Message, "phantom") {
		t.Errorf("unexpected message: %s", diags[0].Message)
	}
	if diags[0].Line != 42 {
		t.Errorf("expected line 42 fallback, got %d", diags[0].Line)
	}
}

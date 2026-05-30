//ff:func feature=validate type=test control=sequence topic=openapi-structural
//ff:what O-6 — dangling required 가 SchemaProperties 인덱스에 있으면 그 줄로 보고

package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestO06_Case14_PropertyLine(t *testing.T) {
	doc := &openapi3.T{Components: &openapi3.Components{Schemas: openapi3.Schemas{
		"Workflow": o06SchemaWithRequired([]string{"id"}, []string{"phantom"}),
	}}}
	fs := &yongol.Fullstack{
		OpenAPIDoc: doc,
		OpenAPILines: &oapiparser.LineIndex{
			Schemas:          map[string]int{"Workflow": 42},
			SchemaProperties: map[string]map[string]int{"Workflow": {"phantom": 99}},
			Paths:            map[string]int{},
			Operations:       map[string]int{},
		},
	}
	diags := o06RequiredPropertyConsistency(fs)
	if len(diags) != 1 {
		t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
	}
	if diags[0].Line != 99 {
		t.Errorf("expected property line 99, got %d", diags[0].Line)
	}
}

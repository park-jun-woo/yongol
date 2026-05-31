//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestBuildResponseField_ZeroCov — $ref / array-of-$ref / 원시 / nil 분기
package ssac

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestBuildResponseField_ZeroCov(t *testing.T) {
	// nil propRef
	if rf := buildResponseField("Name", nil, true); rf.RefType != "" || !rf.IsRequired {
		t.Fatalf("nil propRef: unexpected %+v", rf)
	}
	// direct $ref
	ref := &openapi3.SchemaRef{Ref: "#/components/schemas/Widget"}
	if rf := buildResponseField("widget", ref, false); rf.RefType != "Widget" || rf.IsArray {
		t.Fatalf("direct ref: unexpected %+v", rf)
	}
	// array of $ref
	arr := &openapi3.SchemaRef{Value: &openapi3.Schema{
		Type:  &openapi3.Types{"array"},
		Items: &openapi3.SchemaRef{Ref: "#/components/schemas/Item"},
	}}
	if rf := buildResponseField("items", arr, true); rf.RefType != "Item" || !rf.IsArray {
		t.Fatalf("array ref: unexpected %+v", rf)
	}
	// primitive (no ref, not array)
	prim := &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}}
	if rf := buildResponseField("title", prim, false); rf.RefType != "" || rf.IsArray {
		t.Fatalf("primitive: unexpected %+v", rf)
	}
	// array but items has no ref → no RefType
	arrNoRef := &openapi3.SchemaRef{Value: &openapi3.Schema{
		Type:  &openapi3.Types{"array"},
		Items: &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
	}}
	if rf := buildResponseField("tags", arrNoRef, false); rf.RefType != "" {
		t.Fatalf("array no-ref items: unexpected %+v", rf)
	}
}

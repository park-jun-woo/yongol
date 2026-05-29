//ff:func feature=openapi-parse type=test control=sequence
//ff:what TestExtractResponseArrayItemFields_WithArrayItems — 배열 항목 필드 추출 검증

package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestExtractResponseArrayItemFields_WithArrayItems(t *testing.T) {
	itemSchema := &openapi3.Schema{
		Type: &openapi3.Types{"object"},
		Properties: openapi3.Schemas{
			"id":     &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
			"title":  &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
			"status": &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
		},
	}
	responseSchema := &openapi3.Schema{
		Type: &openapi3.Types{"object"},
		Properties: openapi3.Schemas{
			"items": &openapi3.SchemaRef{
				Value: &openapi3.Schema{
					Type:  &openapi3.Types{"array"},
					Items: &openapi3.SchemaRef{Value: itemSchema},
				},
			},
			"total": &openapi3.SchemaRef{
				Value: &openapi3.Schema{Type: &openapi3.Types{"integer"}},
			},
		},
	}
	doc := &openapi3.T{
		Paths: &openapi3.Paths{},
	}
	doc.Paths.Set("/workflows", &openapi3.PathItem{
		Get: &openapi3.Operation{
			OperationID: "ListWorkflows",
			Responses:   &openapi3.Responses{},
		},
	})
	resp200 := &openapi3.Response{
		Content: openapi3.Content{
			"application/json": &openapi3.MediaType{
				Schema: &openapi3.SchemaRef{Value: responseSchema},
			},
		},
	}
	doc.Paths.Find("/workflows").Get.Responses.Set("200", &openapi3.ResponseRef{Value: resp200})

	result := ExtractResponseArrayItemFields(doc)
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	fields, ok := result["ListWorkflows"]
	if !ok {
		t.Fatal("expected ListWorkflows in result")
	}
	itemFields, ok := fields["items"]
	if !ok {
		t.Fatal("expected 'items' array field")
	}
	if !itemFields["id"] {
		t.Error("expected 'id' in item fields")
	}
	if !itemFields["title"] {
		t.Error("expected 'title' in item fields")
	}
	if !itemFields["status"] {
		t.Error("expected 'status' in item fields")
	}
}

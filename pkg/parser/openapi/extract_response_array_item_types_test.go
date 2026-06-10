//ff:func feature=openapi-parse type=test control=sequence
//ff:what TestExtractResponseArrayItemTypes — 배열 항목 필드 타입 추출·nil doc·비배열/타입없음 분기 검증

package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestExtractResponseArrayItemTypes(t *testing.T) {
	if got := ExtractResponseArrayItemTypes(nil); len(got) != 0 {
		t.Errorf("nil doc: expected empty map, got %v", got)
	}

	itemSchema := &openapi3.Schema{
		Type: &openapi3.Types{"object"},
		Properties: openapi3.Schemas{
			"id":      &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"integer"}}},
			"caption": &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
			"untyped": &openapi3.SchemaRef{Value: &openapi3.Schema{}},
		},
	}
	responseSchema := &openapi3.Schema{
		Type: &openapi3.Types{"object"},
		Properties: openapi3.Schemas{
			"photos": &openapi3.SchemaRef{
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
	doc := &openapi3.T{Paths: &openapi3.Paths{}}
	doc.Paths.Set("/units/{id}", &openapi3.PathItem{
		Get: &openapi3.Operation{OperationID: "GetUnit", Responses: &openapi3.Responses{}},
	})
	doc.Paths.Find("/units/{id}").Get.Responses.Set("200", &openapi3.ResponseRef{
		Value: &openapi3.Response{
			Content: openapi3.Content{
				"application/json": &openapi3.MediaType{
					Schema: &openapi3.SchemaRef{Value: responseSchema},
				},
			},
		},
	})

	result := ExtractResponseArrayItemTypes(doc)
	types, ok := result["GetUnit"]
	if !ok {
		t.Fatalf("expected GetUnit in result: %v", result)
	}
	photoTypes, ok := types["photos"]
	if !ok {
		t.Fatalf("expected 'photos' array field: %v", types)
	}
	if photoTypes["id"] != "integer" || photoTypes["caption"] != "string" {
		t.Errorf("unexpected item types: %v", photoTypes)
	}
	if _, ok := photoTypes["untyped"]; ok {
		t.Errorf("untyped property must be skipped: %v", photoTypes)
	}
	// the non-array "total" property contributes no entry
	if _, ok := types["total"]; ok {
		t.Errorf("non-array field must be absent: %v", types)
	}
}

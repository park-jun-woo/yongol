//ff:func feature=openapi-parse type=test control=iteration dimension=1
//ff:what TestExtractResponseFieldTypes — top-level/object dotted/array item 경로 타입·포맷·allOf·nil/skip 분기 검증

package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestExtractResponseFieldTypes(t *testing.T) {
	if got := ExtractResponseFieldTypes(nil); len(got) != 0 {
		t.Errorf("nil doc: expected empty map, got %v", got)
	}

	responseSchema := &openapi3.Schema{
		AllOf: openapi3.SchemaRefs{
			nil, // skipped
			&openapi3.SchemaRef{Value: &openapi3.Schema{
				Type: &openapi3.Types{"object"},
				Properties: openapi3.Schemas{
					"merged": &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
				},
			}},
		},
		Type: &openapi3.Types{"object"},
		Properties: openapi3.Schemas{
			"nilRef":     nil,                   // skipped
			"nilValue":   &openapi3.SchemaRef{}, // skipped (Value nil)
			"can_delete": &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"boolean"}}},
			"created_at": &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}, Format: "date-time"}},
			"summary": &openapi3.SchemaRef{Value: &openapi3.Schema{
				Type: &openapi3.Types{"object"},
				Properties: openapi3.Schemas{
					"credits_balance": &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"integer"}}},
				},
			}},
			"photos": &openapi3.SchemaRef{Value: &openapi3.Schema{
				Type: &openapi3.Types{"array"},
				Items: &openapi3.SchemaRef{Value: &openapi3.Schema{
					Type: &openapi3.Types{"object"},
					Properties: openapi3.Schemas{
						"url": &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}, Format: "uri"}},
					},
				}},
			}},
		},
	}

	doc := &openapi3.T{Paths: &openapi3.Paths{}}
	doc.Paths.Set("/buildings/{id}", &openapi3.PathItem{
		// no operationId → skipped entirely
		Post: &openapi3.Operation{Responses: &openapi3.Responses{}},
		Get:  &openapi3.Operation{OperationID: "GetBuilding", Responses: &openapi3.Responses{}},
	})
	doc.Paths.Find("/buildings/{id}").Get.Responses.Set("200", &openapi3.ResponseRef{
		Value: &openapi3.Response{
			Content: openapi3.Content{
				"application/json": &openapi3.MediaType{
					Schema: &openapi3.SchemaRef{Value: responseSchema},
				},
			},
		},
	})
	// A non-2xx-only op must contribute nothing.
	doc.Paths.Set("/x", &openapi3.PathItem{
		Delete: func() *openapi3.Operation {
			op := &openapi3.Operation{OperationID: "DelX", Responses: &openapi3.Responses{}}
			op.Responses.Set("404", &openapi3.ResponseRef{Value: &openapi3.Response{
				Content: openapi3.Content{"application/json": &openapi3.MediaType{
					Schema: &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"object"}}},
				}},
			}})
			return op
		}(),
	})

	result := ExtractResponseFieldTypes(doc)
	if _, ok := result["DelX"]; ok {
		t.Errorf("non-2xx-only op must contribute no entry: %v", result["DelX"])
	}
	fields, ok := result["GetBuilding"]
	if !ok {
		t.Fatalf("expected GetBuilding in result: %v", result)
	}

	checks := []struct {
		key, typ, format string
	}{
		{"can_delete", "boolean", ""},
		{"created_at", "string", "date-time"},
		{"summary", "object", ""},
		{"summary.credits_balance", "integer", ""},
		{"photos", "array", ""},
		{"photos.url", "string", "uri"},
		{"merged", "string", ""},
	}
	for _, c := range checks {
		got := fields[c.key]
		if got.Type != c.typ || got.Format != c.format {
			t.Errorf("%s: want %s/%s, got %+v", c.key, c.typ, c.format, got)
		}
	}
	for _, skip := range []string{"nilRef", "nilValue"} {
		if _, ok := fields[skip]; ok {
			t.Errorf("%s must be skipped", skip)
		}
	}
}

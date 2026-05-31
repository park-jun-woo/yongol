//ff:func feature=generate type=test control=sequence
//ff:what TestByName_ZeroCov — generate 폼 액션/필드 해석 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package generate

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestByNameResolveFields_ZeroCov(t *testing.T) {
	schema := openapi3.NewSchema()
	schema.Type = &openapi3.Types{"object"}
	schema.Properties = openapi3.Schemas{
		"Name":  openapi3.NewSchemaRef("", openapi3.NewStringSchema()),
		"Count": openapi3.NewSchemaRef("", openapi3.NewIntegerSchema()),
	}
	reqBody := &openapi3.RequestBodyRef{Value: openapi3.NewRequestBody().WithJSONSchema(schema)}
	op := &openapi3.Operation{OperationID: "CreateItem", RequestBody: reqBody}
	doc := &openapi3.T{Paths: openapi3.NewPaths(
		openapi3.WithPath("/items", &openapi3.PathItem{Post: op}),
	)}

	if fields := resolveFieldsFromOpenAPI(doc, "CreateItem"); len(fields) == 0 {
		t.Errorf("resolveFieldsFromOpenAPI empty")
	}
	// missing op → nil.
	if fields := resolveFieldsFromOpenAPI(doc, "Missing"); fields != nil {
		t.Errorf("resolveFieldsFromOpenAPI missing should be nil")
	}

	ae := actionEntry{opID: "CreateItem", fieldNames: []string{"Name"}}
	if fields := resolveDefaultFields(ae, doc); len(fields) == 0 {
		t.Errorf("resolveDefaultFields from openapi empty")
	}
	// fallback path: unknown op uses fieldNames.
	aeFallback := actionEntry{opID: "Unknown", fieldNames: []string{"A", "B"}}
	if fields := resolveDefaultFields(aeFallback, doc); len(fields) == 0 {
		t.Errorf("resolveDefaultFields fallback empty")
	}
}

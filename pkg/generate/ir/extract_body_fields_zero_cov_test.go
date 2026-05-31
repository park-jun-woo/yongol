//ff:func feature=gen-ir type=test control=sequence
//ff:what ir 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package ir

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestExtractBodyFields_ZeroCov(t *testing.T) {
	schema := openapi3.NewObjectSchema()
	schema.Properties = openapi3.Schemas{
		"name": openapi3.NewSchemaRef("", openapi3.NewStringSchema()),
	}
	schema.Required = []string{"name"}
	op := &openapi3.Operation{
		RequestBody: &openapi3.RequestBodyRef{
			Value: openapi3.NewRequestBody().WithJSONSchema(schema),
		},
	}
	fields := extractBodyFields(op)
	if len(fields) != 1 || fields[0].Name != "name" || !fields[0].Required {
		t.Errorf("body fields wrong: %#v", fields)
	}
	// no request body -> nil
	if got := extractBodyFields(&openapi3.Operation{}); got != nil {
		t.Errorf("no body should be nil")
	}
}

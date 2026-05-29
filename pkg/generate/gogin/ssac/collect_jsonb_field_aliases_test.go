//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what collectJSONBFieldAliases 단위 테스트 (JSONB 프로퍼티만 alias 로 수집)

package ssac

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestCollectJSONBFieldAliases(t *testing.T) {
	schema := &openapi3.Schema{
		Properties: openapi3.Schemas{
			// JSONB: object with additionalProperties unspecified.
			"payload_template": {Value: &openapi3.Schema{Type: &openapi3.Types{"object"}}},
			// non-JSONB scalar.
			"title": {Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
		},
	}
	propNames := []string{"payload_template", "title"}

	got := collectJSONBFieldAliases(schema, propNames)
	if len(got) != 1 {
		t.Fatalf("expected 1 jsonb alias, got %d (%+v)", len(got), got)
	}
	a := got[0]
	if a.jsonName != "payload_template" {
		t.Errorf("jsonName = %q", a.jsonName)
	}
	if a.apiField != "PayloadTemplate" {
		t.Errorf("apiField = %q, want PayloadTemplate", a.apiField)
	}
	if a.dbField != "PayloadTemplate" {
		t.Errorf("dbField = %q, want PayloadTemplate", a.dbField)
	}
	if a.localVar != "payloadTemplateMap" {
		t.Errorf("localVar = %q, want payloadTemplateMap", a.localVar)
	}
}

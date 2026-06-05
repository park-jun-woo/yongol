//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-openapi
//ff:what collectPropTypes — nil 스키마·allOf(nil ref 스킵)·직접 property 수집 분기 검증

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestCollectPropTypes(t *testing.T) {
	// nil schema → no panic, no entries.
	out := map[string]string{}
	collectPropTypes(out, nil)
	if len(out) != 0 {
		t.Fatalf("nil schema: expected empty, got %+v", out)
	}

	// schema with direct props + allOf (incl. a nil ref and a nil-value ref).
	s := &openapi3.Schema{
		Properties: openapi3.Schemas{
			"id": typedSchema("integer"),
		},
		AllOf: openapi3.SchemaRefs{
			nil,          // skipped
			{Value: nil}, // skipped
			{Value: &openapi3.Schema{
				Properties: openapi3.Schemas{
					"title": typedSchema("string"),
					"meta":  typedSchema("object"),
				},
			}},
		},
	}
	out = map[string]string{}
	collectPropTypes(out, s)

	want := map[string]string{
		"id":    "integer",
		"title": "string",
		"meta":  "object",
	}
	if len(out) != len(want) {
		t.Fatalf("collectPropTypes: got %+v, want %+v", out, want)
	}
	for k, v := range want {
		if out[k] != v {
			t.Errorf("out[%q] = %q, want %q", k, out[k], v)
		}
	}
}

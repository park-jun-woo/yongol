//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-openapi
//ff:what collectRequiredNames — 스키마 + allOf required 이름 수집, nil 안전 검증

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestCollectRequiredNames(t *testing.T) {
	s := &openapi3.Schema{
		Required: []string{"a", "b"},
		AllOf: openapi3.SchemaRefs{
			{Value: &openapi3.Schema{Required: []string{"c"}}},
			nil,
		},
	}
	got := collectRequiredNames(s)
	for _, name := range []string{"a", "b", "c"} {
		if _, ok := got[name]; !ok {
			t.Errorf("missing required %q in %+v", name, got)
		}
	}
	if _, ok := got["d"]; ok {
		t.Errorf("unexpected name d")
	}

	if got := collectRequiredNames(nil); len(got) != 0 {
		t.Errorf("nil schema should yield empty, got %+v", got)
	}
}

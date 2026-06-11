//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what allPropsRequired — 전필드 required true / 일부 optional false / 빈 속성 false 검증

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestAllPropsRequired(t *testing.T) {
	all := &openapi3.Schema{
		Properties: openapi3.Schemas{"a": stringProp(), "b": intProp()},
		Required:   []string{"a", "b"},
	}
	if !allPropsRequired(all) {
		t.Errorf("all required should be true")
	}

	partial := &openapi3.Schema{
		Properties: openapi3.Schemas{"a": stringProp(), "b": intProp()},
		Required:   []string{"a"},
	}
	if allPropsRequired(partial) {
		t.Errorf("partial required should be false")
	}

	empty := &openapi3.Schema{Properties: openapi3.Schemas{}}
	if allPropsRequired(empty) {
		t.Errorf("no properties should be false")
	}
}

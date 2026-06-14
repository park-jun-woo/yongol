//ff:func feature=validate type=test control=sequence topic=openapi-ddl
//ff:what resolveFallbackEntity — $ref(DDL-backed)/bare ref/inline 컬럼매칭 순으로 엔티티 폴백 해석

package openapi_ddl

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestResolveFallbackEntity(t *testing.T) {
	idx := idxFor(
		[]ddl.Table{tbl("rules", "id", "name", "created_at")},
		openapi3.Schemas{"Rule": &openapi3.SchemaRef{Value: schemaOf("id", "name")}},
		nil, nil,
	)
	// $ref backed by DDL entity
	if got := resolveFallbackEntity(idx, compRef("Rule", "id", "name")); got != "Rule" {
		t.Errorf("ref Rule → %q, want Rule", got)
	}
	// $ref not DDL-backed
	if got := resolveFallbackEntity(idx, &openapi3.SchemaRef{Ref: "#/components/schemas/Token", Value: schemaOf("a")}); got != "" {
		t.Errorf("ref Token → %q, want \"\"", got)
	}
	// ref without slash exercises the no-slash branch
	if got := resolveFallbackEntity(idx, &openapi3.SchemaRef{Ref: "Rule"}); got != "Rule" {
		t.Errorf("bare ref Rule → %q, want Rule", got)
	}
	// inline → column match fallback
	if got := resolveFallbackEntity(idx, inlineRef("id", "name")); got != "Rule" {
		t.Errorf("inline id,name → %q, want Rule", got)
	}
}

//ff:func feature=validate type=test control=sequence topic=openapi-ddl
//ff:what inferInlineResponseEntity — inline 키 집합으로 유일 DDL 엔티티 추론, <2키/모호/컬럼누락/component 부재는 ""

package openapi_ddl

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestInferInlineResponseEntity(t *testing.T) {
	idx := idxFor(
		[]ddl.Table{
			tbl("rules", "id", "name", "created_at"),
			tbl("notes", "id", "name", "created_at"), // overlaps to force ambiguity
		},
		openapi3.Schemas{
			"Rule": &openapi3.SchemaRef{Value: schemaOf("id", "name")},
			"Note": &openapi3.SchemaRef{Value: schemaOf("id", "name")},
		},
		nil, nil,
	)
	// <2 keys → ""
	if got := inferInlineResponseEntity(idx, []string{"id"}); got != "" {
		t.Errorf("single key → %q, want \"\"", got)
	}
	// ambiguous (both rules and notes contain id,name) → ""
	if got := inferInlineResponseEntity(idx, []string{"id", "name"}); got != "" {
		t.Errorf("ambiguous → %q, want \"\"", got)
	}

	// unique match
	idx2 := idxFor(
		[]ddl.Table{tbl("rules", "id", "name", "created_at")},
		openapi3.Schemas{"Rule": &openapi3.SchemaRef{Value: schemaOf("id", "name")}},
		nil, nil,
	)
	if got := inferInlineResponseEntity(idx2, []string{"id", "name"}); got != "Rule" {
		t.Errorf("unique match → %q, want Rule", got)
	}
	// keys not all present → no candidate
	if got := inferInlineResponseEntity(idx2, []string{"id", "ghost"}); got != "" {
		t.Errorf("missing column → %q, want \"\"", got)
	}

	// table present but no component schema → skipped
	idx3 := idxFor(
		[]ddl.Table{tbl("rules", "id", "name")},
		openapi3.Schemas{}, // no Rule schema
		nil, nil,
	)
	if got := inferInlineResponseEntity(idx3, []string{"id", "name"}); got != "" {
		t.Errorf("no component → %q, want \"\"", got)
	}
}

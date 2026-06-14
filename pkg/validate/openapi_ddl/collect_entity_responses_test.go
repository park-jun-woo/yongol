//ff:func feature=validate type=test control=sequence topic=openapi-ddl
//ff:what collectEntityResponses — nil doc 빈맵, 엔티티별 2xx repr 그룹화, 빈 opID/비엔티티 제외 검증

package openapi_ddl

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestCollectEntityResponses(t *testing.T) {
	// nil OpenAPI doc → empty map
	if got := collectEntityResponses(&yongol.Fullstack{}, &entityIndex{}); len(got) != 0 {
		t.Errorf("nil doc → %v, want empty", got)
	}

	tables := []ddl.Table{tbl("rules", "id", "name", "created_at")}
	schemas := openapi3.Schemas{"Rule": &openapi3.SchemaRef{Value: schemaOf("id", "name", "created_at")}}
	fs := buildCanonicalFS(tables, schemas,
		[]canonOp{
			{"GET", "GetRule", jsonResp("200", compRef("Rule", "id", "name", "created_at"))},
			{"PUT", "UpdateRule", jsonResp("200", inlineRef("id", "name"))},
			// empty opID → skipped
			{"DELETE", "", jsonResp("200", compRef("Rule", "id"))},
			// non-entity (no DDL match) → skipped from grouping
			{"POST", "Login", jsonResp("201", inlineRef("token"))},
		},
		nil, nil,
	)
	idx := buildEntityIndex(fs)
	groups := collectEntityResponses(fs, idx)
	if len(groups["Rule"]) != 2 {
		t.Fatalf("expected 2 Rule reprs, got %d: %+v", len(groups["Rule"]), groups)
	}
	if _, ok := groups[""]; ok {
		t.Error("empty-key group should not exist")
	}
	if len(groups) != 1 {
		t.Errorf("expected only Rule group, got %v", groups)
	}
}

//ff:func feature=validate type=test control=sequence topic=openapi-ddl
//ff:what entityComponent — DDL 테이블 매핑되는 component 만 반환, 빈/비-DDL 모델은 ""

package openapi_ddl

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestEntityComponent(t *testing.T) {
	idx := idxFor(
		[]ddl.Table{tbl("rules", "id", "name")},
		openapi3.Schemas{"Rule": &openapi3.SchemaRef{Value: schemaOf("id", "name")}},
		nil, nil,
	)
	if got := entityComponent(idx, ""); got != "" {
		t.Errorf("empty model → %q, want \"\"", got)
	}
	if got := entityComponent(idx, "Token"); got != "" { // tokens table absent
		t.Errorf("non-DDL model → %q, want \"\"", got)
	}
	if got := entityComponent(idx, "Rule"); got != "Rule" {
		t.Errorf("Rule → %q, want Rule", got)
	}
}

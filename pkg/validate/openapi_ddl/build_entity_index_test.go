//ff:func feature=validate type=test control=sequence topic=openapi-ddl
//ff:what buildEntityIndex — tableExists/schemaForTable/funcByName/Ground 와이어링 검증

package openapi_ddl

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestBuildEntityIndex(t *testing.T) {
	idx := idxFor(
		[]ddl.Table{tbl("rules", "id", "name")},
		openapi3.Schemas{"Rule": &openapi3.SchemaRef{Value: schemaOf("id", "name")}},
		[]ssac.ServiceFunc{{Name: "GetRule"}},
		map[string]string{"SSaC.var.GetRule.rule": "Rule"},
	)
	if !idx.tableExists["rules"] {
		t.Error("tableExists[rules] should be true")
	}
	if idx.schemaForTable["rules"] != "Rule" {
		t.Errorf("schemaForTable[rules] = %q, want Rule", idx.schemaForTable["rules"])
	}
	if idx.funcByName["GetRule"] == nil {
		t.Error("funcByName[GetRule] should be set")
	}
	if idx.g == nil || idx.g.Types["SSaC.var.GetRule.rule"] != "Rule" {
		t.Error("ground types not wired")
	}
}

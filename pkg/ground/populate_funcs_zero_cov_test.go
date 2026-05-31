//ff:func feature=rule type=test control=sequence
//ff:what TestPopulateBatch_ZeroCov — ground populate/helper 함수를 이름으로 직접 호출해 커버 귀속
package ground

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/rego"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestPopulateFuncs_ZeroCov(t *testing.T) {
	fs := &yongol.Fullstack{}
	g := newGround()

	populateSSaC(g, fs)
	populateFunc(g, fs)
	populateSymbolMethods(g, fs)
	populateSSaCRequestUsage(g, fs)
	populateSSaCQueryUsage(g, fs)

	populateDDLColumnFlags(g, ddl.Table{
		Name:        "users",
		Columns:     map[string]ddl.Column{"id": {Name: "id", RawType: "BIGINT", NotNull: true}},
		ColumnOrder: []string{"id"},
	})
	populateDDLDefaults(g, ddl.Table{
		Name:        "users",
		Columns:     map[string]ddl.Column{"id": {Name: "id", RawType: "BIGINT"}},
		ColumnOrder: []string{"id"},
	})

	populateManifestAuth(g, &manifest.Auth{})

	op := &openapi3.Operation{}
	populateOpParams(g, op)
	ops := map[string]*openapi3.Operation{"GetX": op}
	populatePathOps(g, make(rule.StringSet), make(rule.StringSet), ops)
	populatePathOpsParams(g, ops)

	populateRegoPolicy(rego.Policy{}, make(rule.StringSet), make(rule.StringSet), make(rule.StringSet))

	populateResponseFields(g, "Fn", ssac.Sequence{Fields: map[string]string{"a": "x"}})

	applyResponseCodeSchema(g, "GetX", "200", &openapi3.ResponseRef{}, false)
}

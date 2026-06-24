//ff:func feature=validate type=test control=iteration dimension=1
//ff:what kindPresent — SSOT kind별 parse result 존재 여부 판정 검증

package validate

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/parser/features"
	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/rego"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestKindPresent(t *testing.T) {
	cases := []struct {
		name string
		fs   *yongol.Fullstack
		kind yongol.SSOTKind
		want bool
	}{
		{name: "config_nil", fs: &yongol.Fullstack{}, kind: yongol.KindConfig, want: false},
		{name: "config_present", fs: &yongol.Fullstack{Manifest: &manifest.ProjectConfig{}}, kind: yongol.KindConfig, want: true},
		{name: "openapi_nil", fs: &yongol.Fullstack{}, kind: yongol.KindOpenAPI, want: false},
		{name: "openapi_present", fs: &yongol.Fullstack{OpenAPIDoc: &openapi3.T{}}, kind: yongol.KindOpenAPI, want: true},
		{name: "openapi_domain", fs: &yongol.Fullstack{DomainOpenAPIDocs: map[string]*openapi3.T{"public": {}}}, kind: yongol.KindOpenAPI, want: true},
		{name: "openapi_domain_empty", fs: &yongol.Fullstack{DomainOpenAPIDocs: map[string]*openapi3.T{}}, kind: yongol.KindOpenAPI, want: false},
		{name: "ddl_nil", fs: &yongol.Fullstack{}, kind: yongol.KindDDL, want: false},
		{name: "ddl_present", fs: &yongol.Fullstack{DDLTables: []ddl.Table{}}, kind: yongol.KindDDL, want: true},
		{name: "ssac_nil", fs: &yongol.Fullstack{}, kind: yongol.KindSSaC, want: false},
		{name: "ssac_present", fs: &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{}}, kind: yongol.KindSSaC, want: true},
		{name: "states_nil", fs: &yongol.Fullstack{}, kind: yongol.KindStates, want: false},
		{name: "states_present", fs: &yongol.Fullstack{StateDiagrams: []*statemachine.StateDiagram{}}, kind: yongol.KindStates, want: true},
		{name: "policy_nil", fs: &yongol.Fullstack{}, kind: yongol.KindPolicy, want: false},
		{name: "policy_present", fs: &yongol.Fullstack{ParsedPolicies: []rego.Policy{}}, kind: yongol.KindPolicy, want: true},
		{name: "scenario_nil", fs: &yongol.Fullstack{}, kind: yongol.KindScenario, want: false},
		{name: "scenario_entries", fs: &yongol.Fullstack{HurlEntries: []hurl.HurlEntry{}}, kind: yongol.KindScenario, want: true},
		{name: "scenario_files", fs: &yongol.Fullstack{HurlFiles: []string{}}, kind: yongol.KindScenario, want: true},
		{name: "func_nil", fs: &yongol.Fullstack{}, kind: yongol.KindFunc, want: false},
		{name: "func_present", fs: &yongol.Fullstack{ProjectFuncSpecs: []funcspec.FuncSpec{}}, kind: yongol.KindFunc, want: true},
		{name: "func_builtin_only", fs: &yongol.Fullstack{YongolPkgSpecs: []funcspec.FuncSpec{{Package: "auth", Name: "issueToken"}}}, kind: yongol.KindFunc, want: true},
		{name: "stml_nil", fs: &yongol.Fullstack{}, kind: yongol.KindSTML, want: false},
		{name: "stml_present", fs: &yongol.Fullstack{STMLPages: []stml.PageSpec{}}, kind: yongol.KindSTML, want: true},
		{name: "stml_domain", fs: &yongol.Fullstack{DomainSTMLPages: map[string][]stml.PageSpec{"public": {}}}, kind: yongol.KindSTML, want: true},
		{name: "stml_domain_empty", fs: &yongol.Fullstack{DomainSTMLPages: map[string][]stml.PageSpec{}}, kind: yongol.KindSTML, want: false},
		{name: "design_nil", fs: &yongol.Fullstack{}, kind: yongol.KindDesign, want: false},
		{name: "design_present", fs: &yongol.Fullstack{DesignSpec: &design.DesignSpec{}}, kind: yongol.KindDesign, want: true},
		{name: "features_nil", fs: &yongol.Fullstack{}, kind: yongol.KindFeatures, want: false},
		{name: "features_present", fs: &yongol.Fullstack{Features: []features.Feature{}}, kind: yongol.KindFeatures, want: true},
		{name: "unknown_kind", fs: &yongol.Fullstack{}, kind: yongol.SSOTKind("Unknown"), want: false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := kindPresent(c.fs, c.kind)
			if got != c.want {
				t.Errorf("kindPresent(fs, %q) = %v, want %v", c.kind, got, c.want)
			}
		})
	}
}

//ff:func feature=validate type=util control=sequence
//ff:what allSteps — 고정 순서 validation step 목록 반환 (per-SSOT → pair)
package validate

import (
	"github.com/park-jun-woo/yongol/pkg/validate/ddl"
	"github.com/park-jun-woo/yongol/pkg/validate/ddl_rego"
	"github.com/park-jun-woo/yongol/pkg/validate/ddl_statemachine"
	designvalidate "github.com/park-jun-woo/yongol/pkg/validate/design"
	"github.com/park-jun-woo/yongol/pkg/validate/design_manifest"
	"github.com/park-jun-woo/yongol/pkg/validate/domain_security"
	featuresvalidate "github.com/park-jun-woo/yongol/pkg/validate/features"
	"github.com/park-jun-woo/yongol/pkg/validate/features_ddl"
	"github.com/park-jun-woo/yongol/pkg/validate/features_openapi"
	"github.com/park-jun-woo/yongol/pkg/validate/features_statemachine"
	"github.com/park-jun-woo/yongol/pkg/validate/funcspec"
	"github.com/park-jun-woo/yongol/pkg/validate/hurl"
	"github.com/park-jun-woo/yongol/pkg/validate/hurl_manifest"
	"github.com/park-jun-woo/yongol/pkg/validate/hurl_openapi"
	"github.com/park-jun-woo/yongol/pkg/validate/hurl_statemachine"
	"github.com/park-jun-woo/yongol/pkg/validate/initcheck"
	"github.com/park-jun-woo/yongol/pkg/validate/manifest"
	"github.com/park-jun-woo/yongol/pkg/validate/manifest_ddl"
	"github.com/park-jun-woo/yongol/pkg/validate/openapi"
	"github.com/park-jun-woo/yongol/pkg/validate/openapi_ddl"
	"github.com/park-jun-woo/yongol/pkg/validate/openapi_manifest"
	"github.com/park-jun-woo/yongol/pkg/validate/openapi_ssac"
	"github.com/park-jun-woo/yongol/pkg/validate/query"
	"github.com/park-jun-woo/yongol/pkg/validate/query_rego"
	"github.com/park-jun-woo/yongol/pkg/validate/rego"
	"github.com/park-jun-woo/yongol/pkg/validate/rego_manifest"
	"github.com/park-jun-woo/yongol/pkg/validate/ssac"
	"github.com/park-jun-woo/yongol/pkg/validate/ssac_authz"
	"github.com/park-jun-woo/yongol/pkg/validate/ssac_ddl"
	"github.com/park-jun-woo/yongol/pkg/validate/ssac_func"
	"github.com/park-jun-woo/yongol/pkg/validate/ssac_manifest"
	"github.com/park-jun-woo/yongol/pkg/validate/ssac_rego"
	"github.com/park-jun-woo/yongol/pkg/validate/ssac_sqlc"
	"github.com/park-jun-woo/yongol/pkg/validate/ssac_statemachine"
	"github.com/park-jun-woo/yongol/pkg/validate/statemachine"
	"github.com/park-jun-woo/yongol/pkg/validate/stml_design"
	"github.com/park-jun-woo/yongol/pkg/validate/stml_openapi"
	"github.com/park-jun-woo/yongol/pkg/validate/stml_statemachine"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// allSteps returns the validation steps in fixed execution order.
func allSteps() []step {
	return []step{
		{Name: "init", Kinds: nil, Run: initcheck.Run},
		{Name: "manifest", Kinds: []yongol.SSOTKind{yongol.KindConfig}, Run: manifest.Run},
		{Name: "openapi", Kinds: []yongol.SSOTKind{yongol.KindOpenAPI}, Run: openapi.Run},
		{Name: "ddl", Kinds: []yongol.SSOTKind{yongol.KindDDL}, Run: ddl.Run},
		{Name: "query", Kinds: []yongol.SSOTKind{yongol.KindDDL}, Run: query.Run},
		{Name: "ssac", Kinds: []yongol.SSOTKind{yongol.KindSSaC}, Run: ssac.Run},
		{Name: "statemachine", Kinds: []yongol.SSOTKind{yongol.KindStates}, Run: statemachine.Run},
		{Name: "rego", Kinds: []yongol.SSOTKind{yongol.KindPolicy}, Run: rego.Run},
		{Name: "hurl", Kinds: []yongol.SSOTKind{yongol.KindScenario}, Run: hurl.Run},
		{Name: "funcspec", Kinds: []yongol.SSOTKind{yongol.KindFunc}, Run: funcspec.Run},
		{Name: "design", Kinds: []yongol.SSOTKind{yongol.KindDesign}, Run: designvalidate.Run},
		{Name: "features", Kinds: []yongol.SSOTKind{yongol.KindFeatures}, Run: featuresvalidate.Run},
		{Name: "openapi_ddl", Kinds: []yongol.SSOTKind{yongol.KindOpenAPI, yongol.KindDDL}, Run: openapi_ddl.Run},
		{Name: "manifest_ddl", Kinds: []yongol.SSOTKind{yongol.KindConfig, yongol.KindDDL}, Run: manifest_ddl.Run},
		{Name: "openapi_ssac", Kinds: []yongol.SSOTKind{yongol.KindOpenAPI, yongol.KindSSaC}, Run: openapi_ssac.Run},
		{Name: "hurl_openapi", Kinds: []yongol.SSOTKind{yongol.KindScenario, yongol.KindOpenAPI}, Run: hurl_openapi.Run},
		{Name: "hurl_statemachine", Kinds: []yongol.SSOTKind{yongol.KindScenario, yongol.KindStates}, Run: hurl_statemachine.Run},
		{Name: "hurl_manifest", Kinds: []yongol.SSOTKind{yongol.KindScenario, yongol.KindConfig}, Run: hurl_manifest.Run},
		{Name: "openapi_manifest", Kinds: []yongol.SSOTKind{yongol.KindOpenAPI, yongol.KindConfig}, Run: openapi_manifest.Run},
		{Name: "ssac_ddl", Kinds: []yongol.SSOTKind{yongol.KindSSaC, yongol.KindDDL}, Run: ssac_ddl.Run},
		{Name: "ssac_statemachine", Kinds: []yongol.SSOTKind{yongol.KindSSaC, yongol.KindStates}, Run: ssac_statemachine.Run},
		{Name: "ssac_func", Kinds: []yongol.SSOTKind{yongol.KindSSaC, yongol.KindFunc}, Run: ssac_func.Run},
		{Name: "ssac_manifest", Kinds: []yongol.SSOTKind{yongol.KindSSaC, yongol.KindConfig}, Run: ssac_manifest.Run},
		{Name: "ssac_rego", Kinds: []yongol.SSOTKind{yongol.KindSSaC, yongol.KindPolicy}, Run: ssac_rego.Run},
		{Name: "ssac_authz", Kinds: []yongol.SSOTKind{yongol.KindSSaC, yongol.KindPolicy}, Run: ssac_authz.Run},
		{Name: "ssac_sqlc", Kinds: []yongol.SSOTKind{yongol.KindSSaC, yongol.KindDDL}, Run: ssac_sqlc.Run},
		{Name: "ddl_statemachine", Kinds: []yongol.SSOTKind{yongol.KindDDL, yongol.KindStates}, Run: ddl_statemachine.Run},
		{Name: "ddl_rego", Kinds: []yongol.SSOTKind{yongol.KindDDL, yongol.KindPolicy}, Run: ddl_rego.Run},
		{Name: "query_rego", Kinds: []yongol.SSOTKind{yongol.KindDDL, yongol.KindPolicy}, Run: query_rego.Run},
		{Name: "rego_manifest", Kinds: []yongol.SSOTKind{yongol.KindPolicy, yongol.KindConfig}, Run: rego_manifest.Run},
		{Name: "stml_openapi", Kinds: []yongol.SSOTKind{yongol.KindOpenAPI}, Run: stml_openapi.Run},
		{Name: "stml_design", Kinds: []yongol.SSOTKind{yongol.KindSTML, yongol.KindDesign}, Run: stml_design.Run},
		{Name: "stml_statemachine", Kinds: []yongol.SSOTKind{yongol.KindSTML, yongol.KindStates}, Run: stml_statemachine.Run},
		{Name: "domain_security", Kinds: []yongol.SSOTKind{yongol.KindConfig, yongol.KindOpenAPI}, Run: domain_security.Run},
		{Name: "design_manifest", Kinds: []yongol.SSOTKind{yongol.KindConfig}, Run: design_manifest.Run},
		{Name: "features_openapi", Kinds: []yongol.SSOTKind{yongol.KindFeatures, yongol.KindOpenAPI}, Run: features_openapi.Run},
		{Name: "features_ddl", Kinds: []yongol.SSOTKind{yongol.KindFeatures, yongol.KindDDL}, Run: features_ddl.Run},
		{Name: "features_statemachine", Kinds: []yongol.SSOTKind{yongol.KindFeatures, yongol.KindStates}, Run: features_statemachine.Run},
	}
}

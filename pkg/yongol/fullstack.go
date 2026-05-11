//ff:type feature=orchestrator type=model
//ff:what 모든 SSOT 파싱 결과를 담는 풀스택 컨테이너
package yongol

import (
	pg_query "github.com/pganalyze/pg_query_go/v5"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/open-policy-agent/opa/ast"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/parser/rego"
	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/ssacmeta"
)

// Fullstack holds all SSOT parsing results.
// ParseAll() populates this; crosscheck and gen consume it.
type Fullstack struct {
	SpecsDir         string // root directory the SSOTs were parsed from
	Manifest         *manifest.ProjectConfig
	OpenAPIDoc       *openapi3.T
	// OpenAPILines: per-node line index extracted via raw yaml.v3 parse.
	// kin-openapi does not expose line information, so the same openapi.yaml
	// is parsed a second time to build this index. Used to keep file:line
	// accurate in validation diagnostics.
	OpenAPILines     *oapiparser.LineIndex
	DDLResults       []*pg_query.ParseResult
	Policies         []*ast.Module
	ServiceFuncs     []ssac.ServiceFunc
	StateDiagrams    []*statemachine.StateDiagram
	HurlEntries      []hurl.HurlEntry
	ProjectFuncSpecs []funcspec.FuncSpec
	YongolPkgSpecs  []funcspec.FuncSpec
	HurlFiles        []string
	DDLTables        []ddl.Table
	SQLcQueries      []sqlcparser.QuerySpec
	ParsedPolicies   []rego.Policy
	RequestConstraints  map[string]map[string]oapiparser.FieldConstraint
	ResponseConstraints map[string]map[string]oapiparser.FieldConstraint
	STMLPages           []stml.PageSpec
	DesignSpec          *design.DesignSpec
	ParseDiagnostics    []diagnostic.Diagnostic // All errors collected during the parser phase. Gated at the CLI level.
	Presences           map[SSOTKind]SSOTPresence // Presence state (Absent/Declared/Populated) per SSOT kind.
	// SsacInterfaces holds parsed ssac/pkg/*/interface.yaml documents keyed
	// by package name. Populated by parseSsacInterfaces during ParseAll and
	// consumed by Phase002 codegen (pkg/generate/gogin/infra/) plus Phase004/005
	// validate rules. Absent key ⇒ package declares no DB ports (or is not a
	// DB-using ssac package).
	SsacInterfaces      map[string]*ssacmeta.PackageInterface
	ground              *rule.Ground
}

//ff:type feature=orchestrator type=model
//ff:what 모든 SSOT 파싱 결과를 담는 풀스택 컨테이너
package yongol

import (
	pg_query "github.com/pganalyze/pg_query_go/v5"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/open-policy-agent/opa/ast"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/parser/features"
	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/parser/rego"
	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/ssacmeta"
)

// Fullstack holds all SSOT parsing results.
// ParseAll() populates this; crosscheck and gen consume it.
type Fullstack struct {
	SpecsDir   string // root directory the SSOTs were parsed from
	Manifest   *manifest.ProjectConfig
	OpenAPIDoc *openapi3.T
	// OpenAPILines: per-node line index extracted via raw yaml.v3 parse.
	// kin-openapi does not expose line information, so the same openapi.yaml
	// is parsed a second time to build this index. Used to keep file:line
	// accurate in validation diagnostics.
	OpenAPILines        *oapiparser.LineIndex
	DDLResults          []*pg_query.ParseResult
	Policies            []*ast.Module
	ServiceFuncs        []ssac.ServiceFunc
	StateDiagrams       []*statemachine.StateDiagram
	HurlEntries         []hurl.HurlEntry
	ProjectFuncSpecs    []funcspec.FuncSpec
	YongolPkgSpecs      []funcspec.FuncSpec
	HurlFiles           []string
	DDLTables           []ddl.Table
	SQLcQueries         []sqlcparser.QuerySpec
	ParsedPolicies      []rego.Policy
	RequestConstraints  map[string]map[string]oapiparser.FieldConstraint
	ResponseConstraints map[string]map[string]oapiparser.FieldConstraint
	STMLPages           []stml.PageSpec
	Layouts             []stml.LayoutSpec
	// Sitemap is the parsed frontend/sitemap.html — the central site-structure
	// declaration (plans/stml/sitemap Phase001). nil = file absent: every
	// sitemap-derived behavior stays off (optional-file backward compatibility).
	Sitemap          *stml.SitemapSpec
	Features         []features.Feature
	FeatureTables    map[string]features.TableDef
	DesignSpec       *design.DesignSpec
	ParseDiagnostics []diagnostic.Diagnostic   // All errors collected during the parser phase. Gated at the CLI level.
	Presences        map[SSOTKind]SSOTPresence // Presence state (Absent/Declared/Populated) per SSOT kind.
	// DomainPresences holds per-domain SSOT presence for multi-domain projects:
	// domain name → SSOT kind → presence. Only OpenAPI and STML are split per
	// domain (everything else is shared). nil for single-site projects; populated
	// by ParseAll's domain loop (Phase004) when fs.Manifest.Domains is non-empty.
	DomainPresences map[string]map[SSOTKind]SSOTPresence
	// Per-domain SSOT data for multi-domain projects (domains: declared). Keyed by
	// domain name. nil for single-site projects (the singular fields above are used
	// instead). Populated by ParseAll's domain loop (Phase004); consumed via the
	// helper accessors (IsDomained/DomainNames/AllOpenAPIDocs/AllSTMLPages/DomainView).
	DomainOpenAPIDocs map[string]*openapi3.T
	// DomainOpenAPILines is declared here (Phase003 data model) but populated later
	// in Phase004's ParseAll domain loop — same per-node line index as the singular
	// OpenAPILines, one per domain.
	DomainOpenAPILines map[string]*oapiparser.LineIndex
	DomainSTMLPages    map[string][]stml.PageSpec
	DomainSitemaps     map[string]*stml.SitemapSpec
	DomainLayouts      map[string][]stml.LayoutSpec
	// SsacInterfaces holds parsed ssac/pkg/*/interface.yaml documents keyed
	// by package name. Populated by parseSsacInterfaces during ParseAll and
	// consumed by Phase002 codegen (pkg/generate/gogin/infra/) plus Phase004/005
	// validate rules. Absent key ⇒ package declares no DB ports (or is not a
	// DB-using ssac package).
	SsacInterfaces map[string]*ssacmeta.PackageInterface
	// FuncPackageTypes holds all struct types collected from the func spec
	// directory tree, keyed by package name → type name → fields. Populated
	// by parseFuncIfPresent via funcspec.CollectAllPackageTypes and consumed
	// by the ssac codegen emitter to resolve inner (non-response) types
	// referenced from @call result structs (BUG-149 / Phase006).
	FuncPackageTypes map[string]map[string][]funcspec.Field
	ground           *rule.Ground
}

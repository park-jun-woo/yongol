//ff:type feature=validate type=model
//ff:what step — 단일 검증 단위 (SSOT 또는 pair 폴더) 메타데이터
package validate

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// step is a single validation unit — one SSOT folder or one pair folder.
// Run receives the fully wired Fullstack (including fs.Ground()); no separate
// ground parameter needed.
type step struct {
	Name  string
	Kinds []yongol.SSOTKind
	Run   func(fs *yongol.Fullstack) []diagnostic.Diagnostic

	// DomainAware marks a step whose Run body is already multi-domain aware —
	// it iterates fs.AllOpenAPIDocs()/fs.AllSTMLPages() or the merged Ground,
	// or guards on fs.Manifest.Domains (Phase005: stml_openapi, openapi_ssac,
	// domain_security, features_openapi). In domain mode it runs ONCE on the
	// full fs so cross-domain detection (e.g. XDO-90) and the merged coverage
	// sets stay intact; running it per-view would double-run or lose that.
	DomainAware bool

	// DomainMerged marks a single-doc OpenAPI step whose rule bodies read the
	// singular fs.OpenAPIDoc AND include reverse rules that consult the FULL
	// operationId / securityScheme set against global manifest/hurl config
	// (openapi_manifest SEC-05, hurl_openapi route coverage). Per-domain-view
	// would false-positive (e.g. the rate_limit op Login lives only in the
	// public doc → flagged in the admin view). In domain mode it runs ONCE on
	// fs.MergedOpenAPIView() so those rules see every domain at once, with zero
	// rule-body changes. Mutually exclusive with DomainAware.
	DomainMerged bool
}

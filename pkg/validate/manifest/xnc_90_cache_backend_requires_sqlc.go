//ff:func feature=validate type=rule control=sequence topic=manifest-infra
//ff:what XNC-90 — manifest.cache.backend=postgres 시 canonical DDL + sqlc 쿼리 존재 강제

package manifest

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xnc90CacheBackendRequiresSQLC validates XNC-90: when the user opts into
// manifest.cache.backend == "postgres", the companion DDL table
// `fullend_cache` and sqlc queries `CacheSet / CacheGet / CacheDelete`
// (declared as ports in ssac/pkg/cache/interface.yaml) must exist in the
// user-authored SSOTs. Missing entries surface as a single ERROR so the
// user sees every required entity in one advice blob.
//
// Advice payload: the corresponding ssac interface.yaml's canonical_ddl +
// canonical_queries blocks concatenated verbatim. This keeps the emitter
// catalog-free — renaming a port in ssac automatically propagates to the
// advice text the next time yongol is rebuilt.
func xnc90CacheBackendRequiresSQLC(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	return validateBuiltinBackend(fs, backendSpec{
		Pkg:        "cache",
		Cfg:        cacheCfg(fs),
		RequireDDL: "fullend_cache",
		RequireQueries: []string{
			"CacheSet", "CacheGet", "CacheDelete",
		},
		RuleID: "XNC-90",
	})
}

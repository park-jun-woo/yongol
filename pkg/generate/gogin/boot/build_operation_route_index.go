//ff:func feature=gen-gogin type=util control=sequence topic=dos-guard
//ff:what buildOperationRouteIndex — operationId → "METHOD <prefix><path>" 매핑 (도메인 인지)

package boot

import "github.com/park-jun-woo/yongol/pkg/yongol"

// buildOperationRouteIndex maps each operationId to its gin route key
// ("METHOD <route>"). The route key must match the runtime FullPath the
// rate-limit / body-limit middlewares compare against.
//
// Domain mode: each domain's OpenAPI paths are RELATIVE and mounted under its
// route_prefix (r.Group(cfg.RoutePrefix) + RegisterHandlers), so keys are
// prefixed per domain and merged into one map (operationIds globally unique,
// XDO-90). Single-site: the singular OpenAPIDoc is indexed with no prefix —
// byte-identical to the pre-domain behaviour. Returns an empty map when no
// document is available so callers can iterate safely.
func buildOperationRouteIndex(fs *yongol.Fullstack) map[string]string {
	idx := map[string]string{}
	if fs == nil {
		return idx
	}
	if fs.IsDomained() {
		for _, name := range fs.DomainNames() {
			indexOpenAPIDoc(idx, fs.DomainOpenAPIDocs[name], fs.Manifest.Domains[name].RoutePrefix)
		}
		return idx
	}
	indexOpenAPIDoc(idx, fs.OpenAPIDoc, "")
	return idx
}

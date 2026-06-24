//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what bearerDomainPrefixes — bearer 전용 도메인의 route_prefix 목록 (CSRF 면제 경로용)

package boot

import "github.com/park-jun-woo/yongol/pkg/yongol"

// bearerDomainPrefixes returns the route_prefix of every domain whose resolved
// auth_mode is bearer (Phase008 §4). The global Csrf middleware exempts these
// prefixes — bearer transport is CSRF-immune (Authorization headers are not
// auto-sent cross-origin) — while cookie/hybrid domain paths stay verified.
// Returns nil for single-site projects.
func bearerDomainPrefixes(fs *yongol.Fullstack) []string {
	if fs == nil || !fs.IsDomained() {
		return nil
	}
	var out []string
	for _, name := range fs.DomainNames() {
		if domainWireMode(fs, name) == "bearer" {
			out = append(out, fs.Manifest.Domains[name].RoutePrefix)
		}
	}
	return out
}

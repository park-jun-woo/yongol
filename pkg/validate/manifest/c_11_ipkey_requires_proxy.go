//ff:func feature=validate type=rule control=iteration dimension=1 topic=manifest-structural
//ff:what C-11 — key=ip(또는 미설정→default ip) rate_limit + trusted_proxies 미설정 → WARNING (BUG-117)

package manifest

import (
	"sort"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// c11IpKeyRequiresProxy warns when an IP-keyed rate_limit entry exists but
// backend.http.trusted_proxies is unset.
//
// BUG-117 / Phase-C2 — after Phase-C1 the generated router calls
// SetTrustedProxies(nil) by default, so c.ClientIP() uses RemoteAddr only.
// That is the safe default for directly exposed servers, but behind a
// reverse proxy RemoteAddr is always the proxy address: every client then
// collapses onto one limiter key and the IP-keyed rate limit degenerates
// into a global throttle (combining BUG-115's limiter with BUG-117's
// proxy-trust gap).
//
// Validate cannot know the deployment topology (behind a proxy vs directly
// exposed), so this is a WARNING, not an ERROR: directly exposed
// deployments are correct as-is and may ignore it. An entry with an empty
// key counts as "ip" because codegen defaults the key axis to ip
// (block_rate_limit.go). When trusted_proxies is declared (non-empty) the
// operator has made the proxy decision explicitly and the rule stays
// silent — env-time overrides via BACKEND_HTTP_TRUSTED_PROXIES remain an
// operational concern outside static validation.
func c11IpKeyRequiresProxy(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || fs.Manifest == nil {
		return nil
	}
	b := fs.Manifest.Backend
	if b.HTTP != nil && len(b.HTTP.TrustedProxies) > 0 {
		// Proxy trust is declared; c.ClientIP() resolves the real client
		// behind the listed CIDRs.
		return nil
	}

	// Collect IP-keyed entries in deterministic order for a stable message.
	var ipKeyed []string
	for opID, entry := range b.RateLimit {
		if entry.Key == "" || entry.Key == "ip" {
			ipKeyed = append(ipKeyed, opID)
		}
	}
	if len(ipKeyed) == 0 {
		return nil
	}
	sort.Strings(ipKeyed)

	return []diagnostic.Diagnostic{{
		File:  "manifest.yaml",
		Phase: diagnostic.PhaseValidate,
		Level: diagnostic.LevelWarning,
		Message: "[C-11] backend.rate_limit entries keyed by ip (" +
			strings.Join(ipKeyed, ", ") +
			") with backend.http.trusted_proxies unset — behind a reverse proxy " +
			"c.ClientIP() always returns the proxy address, so the IP-keyed rate " +
			"limit collapses every client onto one key (effectively a global limit)",
		Advice: "If this backend is deployed behind a reverse proxy, declare the " +
			"proxy CIDR ranges in backend.http.trusted_proxies (e.g. " +
			"[\"10.0.0.0/8\"]) so IP-keyed rate limits see the real client IP " +
			"(BACKEND_HTTP_TRUSTED_PROXIES can override at runtime). If the " +
			"server is directly exposed, the current setting is correct and " +
			"this warning can be ignored",
	}}
}

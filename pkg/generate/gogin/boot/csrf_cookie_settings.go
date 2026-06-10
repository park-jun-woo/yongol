//ff:func feature=gen-gogin type=util control=sequence topic=csrf
//ff:what csrfCookieSettings — CsrfConfig 기본값 적용 후 생성에 필요한 값들 추출

package boot

import "github.com/park-jun-woo/yongol/pkg/parser/manifest"

// csrfCookieSettings extracts the CSRF cookie/header/exempt/max-age/
// secure values with the manifest-resolved defaults applied
// (CsrfConfig.ResolvedCookieName / ResolvedHeaderName / ResolvedMaxAge —
// the single source of the CSRF defaults). Secure is always true — the
// knob was removed in Phase020 to avoid accidentally shipping an
// insecure cookie on production manifests.
func csrfCookieSettings(c *manifest.CsrfConfig) (string, string, []string, int, bool) {
	cookieName := c.ResolvedCookieName()
	headerName := c.ResolvedHeaderName()
	maxAge := c.ResolvedMaxAge()
	secure := true
	var exempt []string
	if c != nil {
		exempt = c.ExemptPaths
	}
	return cookieName, headerName, exempt, maxAge, secure
}

//ff:func feature=gen-gogin type=util control=sequence topic=csrf
//ff:what csrfCookieSettings — CsrfConfig 기본값 적용 후 생성에 필요한 값들 추출

package boot

import "github.com/park-jun-woo/yongol/pkg/parser/manifest"

// csrfCookieSettings extracts the CSRF cookie/header/exempt/max-age/
// secure values with Phase005 defaults applied. Secure is always true —
// the knob was removed in Phase020 to avoid accidentally shipping an
// insecure cookie on production manifests.
func csrfCookieSettings(c *manifest.CsrfConfig) (string, string, []string, int, bool) {
	cookieName := "XSRF-TOKEN"
	headerName := "X-XSRF-TOKEN"
	var exempt []string
	maxAge := 86400
	secure := true
	if c == nil {
		return cookieName, headerName, exempt, maxAge, secure
	}
	if c.CookieName != "" {
		cookieName = c.CookieName
	}
	if c.HeaderName != "" {
		headerName = c.HeaderName
	}
	exempt = c.ExemptPaths
	if c.MaxAge > 0 {
		maxAge = c.MaxAge
	}
	return cookieName, headerName, exempt, maxAge, secure
}

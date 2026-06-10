//ff:func feature=projectconfig type=util control=sequence
//ff:what ResolvedCookieName — csrf.cookie_name 해석값 (nil/빈 값 → "XSRF-TOKEN")

package manifest

// ResolvedCookieName returns the effective csrf.cookie_name ("XSRF-TOKEN"
// when unset). Prefer this over reading CookieName directly. The same
// default constant is mirrored by the self-contained runtime fallback in
// pkg/generate/gogin/middleware/csrf_source.go — keep both in sync.
func (c *CsrfConfig) ResolvedCookieName() string {
	if c == nil || c.CookieName == "" {
		return "XSRF-TOKEN"
	}
	return c.CookieName
}

//ff:func feature=projectconfig type=util control=sequence
//ff:what ResolvedHeaderName — csrf.header_name 해석값 (nil/빈 값 → "X-XSRF-TOKEN")

package manifest

// ResolvedHeaderName returns the effective csrf.header_name ("X-XSRF-TOKEN"
// when unset). Prefer this over reading HeaderName directly. The same
// default constant is mirrored by the self-contained runtime fallback in
// pkg/generate/gogin/middleware/csrf_source.go — keep both in sync.
func (c *CsrfConfig) ResolvedHeaderName() string {
	if c == nil || c.HeaderName == "" {
		return "X-XSRF-TOKEN"
	}
	return c.HeaderName
}

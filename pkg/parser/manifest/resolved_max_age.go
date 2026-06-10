//ff:func feature=projectconfig type=util control=sequence
//ff:what ResolvedMaxAge — csrf.max_age 해석값 (nil/비양수 → 86400)

package manifest

// ResolvedMaxAge returns the effective csrf.max_age in seconds (86400
// when unset or non-positive). Prefer this over reading MaxAge directly.
// The same default constant is mirrored by the self-contained runtime
// fallback in pkg/generate/gogin/middleware/csrf_source.go.
func (c *CsrfConfig) ResolvedMaxAge() int {
	if c == nil || c.MaxAge <= 0 {
		return 86400
	}
	return c.MaxAge
}

//ff:func feature=manifest type=test control=sequence
//ff:what ResolvedCookieName/HeaderName/MaxAge — nil/빈 값은 기본값, 명시 값은 그대로

package manifest

import (
	"testing"
)

func TestResolvedCsrf(t *testing.T) {
	var nilCsrf *CsrfConfig
	if got := nilCsrf.ResolvedCookieName(); got != "XSRF-TOKEN" {
		t.Errorf("nil csrf cookie → %q, want XSRF-TOKEN", got)
	}
	if got := nilCsrf.ResolvedHeaderName(); got != "X-XSRF-TOKEN" {
		t.Errorf("nil csrf header → %q, want X-XSRF-TOKEN", got)
	}
	if got := nilCsrf.ResolvedMaxAge(); got != 86400 {
		t.Errorf("nil csrf max_age → %d, want 86400", got)
	}
	if got := (&CsrfConfig{}).ResolvedCookieName(); got != "XSRF-TOKEN" {
		t.Errorf("empty cookie → %q, want XSRF-TOKEN", got)
	}
	if got := (&CsrfConfig{}).ResolvedHeaderName(); got != "X-XSRF-TOKEN" {
		t.Errorf("empty header → %q, want X-XSRF-TOKEN", got)
	}
	if got := (&CsrfConfig{}).ResolvedMaxAge(); got != 86400 {
		t.Errorf("empty max_age → %d, want 86400", got)
	}
	override := &CsrfConfig{CookieName: "MY-XSRF", HeaderName: "X-My-CSRF", MaxAge: 120}
	if got := override.ResolvedCookieName(); got != "MY-XSRF" {
		t.Errorf("explicit cookie → %q, want MY-XSRF", got)
	}
	if got := override.ResolvedHeaderName(); got != "X-My-CSRF" {
		t.Errorf("explicit header → %q, want X-My-CSRF", got)
	}
	if got := override.ResolvedMaxAge(); got != 120 {
		t.Errorf("explicit max_age → %d, want 120", got)
	}
}

//ff:func feature=validate type=rule control=sequence topic=manifest-structural
//ff:what C-8 — backend.rate_limit 에 Login 항목 필수 (auth 의 핵심 엔드포인트)

package manifest

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// c08RateLimitLoginRequired enforces that rate_limit contains at least a
// "Login" entry. Login is the primary brute-force attack surface — omitting
// it defeats the purpose of having rate limiting on auth endpoints.
func c08RateLimitLoginRequired(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || fs.Manifest == nil {
		return nil
	}
	if fs.Manifest.Backend.Auth == nil {
		return nil
	}
	if len(fs.Manifest.Backend.RateLimit) == 0 {
		// C-7 already flags the missing rate_limit section.
		return nil
	}
	if _, ok := fs.Manifest.Backend.RateLimit["Login"]; ok {
		return nil
	}
	return []diagnostic.Diagnostic{{
		File:    "manifest.yaml",
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: "[C-8] backend.rate_limit requires a Login entry",
		Advice: "Add a Login entry to backend.rate_limit. Login is the primary " +
			"brute-force target and must be rate-limited. Example:\n" +
			"  backend:\n" +
			"    rate_limit:\n" +
			"      Login:\n" +
			"        rate: 5\n" +
			"        period: \"1m\"\n" +
			"        key: ip",
	}}
}

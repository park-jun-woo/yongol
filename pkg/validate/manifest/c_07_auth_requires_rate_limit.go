//ff:func feature=validate type=rule control=sequence topic=manifest-structural
//ff:what C-7 — backend.auth 존재 시 backend.rate_limit 필수 (brute-force 방어 필수)

package manifest

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// c07AuthRequiresRateLimit enforces that every manifest declaring
// backend.auth also declares backend.rate_limit. Without rate limiting,
// authentication endpoints are exposed to brute-force and credential-
// stuffing attacks. Failing fast at validate time prevents the gap from
// reaching production.
func c07AuthRequiresRateLimit(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || fs.Manifest == nil {
		return nil
	}
	if fs.Manifest.Backend.Auth == nil {
		return nil
	}
	if len(fs.Manifest.Backend.RateLimit) > 0 {
		return nil
	}
	return []diagnostic.Diagnostic{{
		File:    "manifest.yaml",
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: "[C-7] backend.auth requires backend.rate_limit — brute-force defense is mandatory",
		Advice: "Add a backend.rate_limit section to manifest.yaml with at least " +
			"a Login entry. Example:\n" +
			"  backend:\n" +
			"    rate_limit:\n" +
			"      Login:\n" +
			"        rate: 5\n" +
			"        period: \"1m\"\n" +
			"        key: ip",
	}}
}

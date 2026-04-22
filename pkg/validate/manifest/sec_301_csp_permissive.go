//ff:func feature=validate type=rule control=sequence topic=manifest-security-headers
//ff:what SEC-301 — CSP default-src 가 '*' 또는 'unsafe-eval' 을 허용하면 WARNING

package manifest

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// sec301CspPermissive flags overly permissive CSP default-src configurations.
// Allowing "*" or "'unsafe-eval'" in default-src defeats the primary purpose
// of CSP (blocking arbitrary script origins and eval-based injections).
// Levels as WARNING because some legitimate dashboards ship with these
// relaxations during early rollout — operators see the risk but codegen
// still succeeds.
func sec301CspPermissive(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || fs.Manifest == nil {
		return nil
	}
	sh := fs.Manifest.Backend.SecurityHeaders
	if sh == nil || sh.CSP == nil {
		return nil
	}
	srcs := sh.CSP.Directives["default-src"]
	if len(srcs) == 0 {
		return nil
	}
	var offenders []string
	for _, s := range srcs {
		if s == "*" || s == "'unsafe-eval'" {
			offenders = append(offenders, s)
		}
	}
	if len(offenders) == 0 {
		return nil
	}
	return []diagnostic.Diagnostic{{
		File:    "manifest.yaml",
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelWarning,
		Message: "[SEC-301] backend.security_headers.csp.default-src 에 " + quoted(offenders[0]) + " 가 포함되어 CSP 보호 효과가 약화됩니다",
		Advice:  "default-src 에서 '*' / 'unsafe-eval' 을 제거하고 필요한 origin 만 허용하세요 (예: default-src 'self')",
	}}
}

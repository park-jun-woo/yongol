//ff:func feature=gen-ir type=util control=sequence
//ff:what securityHeadersEnabled -- manifest.backend.security_headers.enabled 여부 (기본 true)

package ir

import "github.com/park-jun-woo/yongol/pkg/yongol"

// securityHeadersEnabled returns true when security headers are active.
// Defaults to true (opt-out).
func securityHeadersEnabled(fs *yongol.Fullstack) bool {
	if fs == nil || fs.Manifest == nil {
		return true
	}
	sh := fs.Manifest.Backend.SecurityHeaders
	if sh == nil {
		return true
	}
	if sh.Enabled == nil {
		return true
	}
	return *sh.Enabled
}

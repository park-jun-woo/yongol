//ff:func feature=gen-gogin type=util control=sequence topic=security-headers
//ff:what hasSecurityHeaders — manifest.backend.security_headers.enabled (기본 true) 여부

package boot

import "github.com/park-jun-woo/yongol/pkg/yongol"

// hasSecurityHeaders returns true when security headers should be wired into
// the generated main.go. Defaults to true (opt-out): a missing manifest
// block or missing enabled flag still yields true so every project ships
// with browser security headers out of the box.
func hasSecurityHeaders(fs *yongol.Fullstack) bool {
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

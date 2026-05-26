//ff:func feature=gen-ir type=generator control=sequence
//ff:what buildSecurityHeadersConfig -- manifest.backend.security_headers → SecurityHeadersConfig 변환

package ir

import "github.com/park-jun-woo/yongol/pkg/yongol"

// buildSecurityHeadersConfig returns nil when explicitly disabled.
func buildSecurityHeadersConfig(fs *yongol.Fullstack) *SecurityHeadersConfig {
	if !securityHeadersEnabled(fs) {
		return nil
	}
	cfg := &SecurityHeadersConfig{
		Profile:         "production",
		HSTSMaxAge:      31536000,
		HSTSIncludeSubs: true,
		HSTSPreload:     false,
		CSPEnabled:      true,
		CSPReportOnly:   false,
		XFrameOptions:   "DENY",
		ReferrerPolicy:  "strict-origin-when-cross-origin",
	}
	if fs == nil || fs.Manifest == nil {
		return cfg
	}
	sh := fs.Manifest.Backend.SecurityHeaders
	if sh == nil {
		return cfg
	}
	applySecurityHeadersOverrides(cfg, sh)
	return cfg
}

//ff:func feature=gen-ir type=util control=selection
//ff:what applySecurityHeadersOverrides -- manifest 보안 헤더 오버라이드 적용

package ir

import "github.com/park-jun-woo/yongol/pkg/parser/manifest"

// applySecurityHeadersOverrides applies manifest overrides to the
// security headers config.
func applySecurityHeadersOverrides(cfg *SecurityHeadersConfig, sh *manifest.SecurityHeadersConfig) {
	if sh.Profile != "" {
		cfg.Profile = sh.Profile
	}
	if sh.HSTS != nil {
		if sh.HSTS.MaxAge > 0 {
			cfg.HSTSMaxAge = sh.HSTS.MaxAge
		}
		cfg.HSTSIncludeSubs = sh.HSTS.IncludeSubDomains
		cfg.HSTSPreload = sh.HSTS.Preload
	}
	if sh.CSP != nil {
		cfg.CSPReportOnly = sh.CSP.ReportOnly
		if len(sh.CSP.Directives) > 0 {
			cfg.CSPDirectives = sh.CSP.Directives
		}
	}
	if sh.XFrameOptions != "" {
		cfg.XFrameOptions = sh.XFrameOptions
	}
	if sh.ReferrerPolicy != "" {
		cfg.ReferrerPolicy = sh.ReferrerPolicy
	}
	if len(sh.PermissionsPolicy) > 0 {
		cfg.PermissionsPolicy = sh.PermissionsPolicy
	}
	// Apply profile-specific defaults
	switch cfg.Profile {
	case "dev":
		cfg.CSPReportOnly = true
	case "api":
		cfg.CSPEnabled = false
	}
}
